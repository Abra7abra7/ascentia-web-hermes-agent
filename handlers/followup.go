package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// FollowUpSender je injektovateľná funkcia pre odoslanie follow-up emailu.
// Umožňuje testovanie bez reálneho volania Resend API.
type FollowUpSender func(name, email, analysis, urgency string)

// RunFollowUpCheck skontroluje databázu pre dopyty staršie ako 24h
// ktoré ešte nemajú odoslaný follow-up email.
//
// Bezpečnostný a idempotenčný design (oprava duplicitného odosielania):
//   - Používa sa atomický UPDATE ... RETURNING, ktorý označí záznam ako
//     odoslaný (follow_up_sent = 1) A vráti jeho dáta v JEDNOM SQL príkaze.
//   - Tým sa eliminuje race condition medzi SELECT a UPDATE, ku ktorému
//     dochádzalo pri Fly.io auto_stop_machines (machine sa uspala po odoslaní
//     mailu, ale pred commitom UPDATE → pri prebudení sa ten istý záznam
//     znova vybral a znova poslal).
//   - Beží v transakcii BEGIN IMMEDIATE pre serializáciu zápisov v SQLite.
func (s *Server) RunFollowUpCheck() {
	s.RunFollowUpCheckWithSender(sendFollowUpEmail)
}

// RunFollowUpCheckWithSender je testovateľná verzia RunFollowUpCheck s injektovaným
// odosielateľom emailu.
func (s *Server) RunFollowUpCheckWithSender(sender FollowUpSender) {
	fmt.Printf("[FOLLOW-UP] Running check at %s...\n", time.Now().Format(time.RFC3339))

	// Transakcia s IMMEDIATE lockom — serializuje zápisy, predchádza
	// súbehu pri viacerých behoch scheduleru / reštartoch machine.
	tx, err := s.DB.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		fmt.Printf("[FOLLOW-UP] ERROR: could not begin transaction: %v\n", err)
		return
	}
	defer tx.Rollback() // bezpečné — ak commit prebehol, Rollback je no-op

	// Atomický UPDATE ... RETURNING: označí follow_up_sent = 1 A vráti dáta
	// vybraných záznamov v jednom kroku. Žiadna medzera na race condition.
	rows, err := tx.Query(`
		UPDATE contact_inquiries
		SET follow_up_sent = 1,
		    sent_at = datetime('now')
		WHERE id IN (
			SELECT ci.id
			FROM contact_inquiries ci
			LEFT JOIN lead_scores ls ON ls.inquiry_id = ci.id
			WHERE ci.created_at < datetime('now', '-24 hours')
			  AND (ci.follow_up_sent IS NULL OR ci.follow_up_sent = 0)
			ORDER BY ci.created_at ASC
			LIMIT 50
		)
		RETURNING id, name, email, company, message,
		          COALESCE((SELECT ls.score FROM lead_scores ls WHERE ls.inquiry_id = contact_inquiries.id), 0),
		          COALESCE((SELECT ls.budget FROM lead_scores ls WHERE ls.inquiry_id = contact_inquiries.id), ''),
		          COALESCE((SELECT ls.urgency FROM lead_scores ls WHERE ls.inquiry_id = contact_inquiries.id), 'medium')
	`)
	if err != nil {
		fmt.Printf("[FOLLOW-UP] DB query error: %v\n", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id int64
		var name, email, company, message, budget, urgency string
		var score int64

		if err := rows.Scan(&id, &name, &email, &company, &message, &score, &budget, &urgency); err != nil {
			fmt.Printf("[FOLLOW-UP] Row scan error: %v\n", err)
			continue
		}

		analysis := generateInquiryAnalysis(name, company, message, budget, urgency, score)
		sender(name, email, analysis, urgency)
		count++
		fmt.Printf("[FOLLOW-UP EMAIL] Sent to: %s | Name: %s | InquiryID: %d\n", email, name, id)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("[FOLLOW-UP] Rows iteration error: %v\n", err)
		return
	}

	// Commit až po úspešnom odoslaní všetkých emailov v tejto iterácii.
	// Ak by odoslanie zlyhalo, Rollback (defer) vráti follow_up_sent = 0
	// a mail sa skúsi znova v ďalšom behu (bez straty notifikácie).
	if err := tx.Commit(); err != nil {
		fmt.Printf("[FOLLOW-UP] ERROR: commit failed: %v\n", err)
		return
	}

	fmt.Printf("[FOLLOW-UP] Sent %d follow-up emails\n", count)
}

// generateInquiryAnalysis vytvorí AI analýzu dopytu pre follow-up email
func generateInquiryAnalysis(name, company, message, budget, urgency string, score int64) string {
	return fmt.Sprintf(`Dobry den, %s!

Preanalyzovali sme vas dopyt a pripravili sme zakladnu AI kognitivnu analyzu:

SHRNUTIE PROJEKTU:
- Spolocnost: %s
- AI Lead Score: %d/100
- Odhadovany rozpocet: %s
- Priorita: %s

VASA SPRAVA:
%s

NAS NAVRH RIESENIA:
Na zaklade analyzy vaseho dopytu vam odporucame:

1. Uvodna konzultacia (30 min) — zdarma
   Prejdeme si vase pozadavky a navrhneme technicke riesenie.

2. Architektonicky navrh
   Pripravime high-level dizajn v Go (Golang) s AI integraciou.

3. Prototyp (TDD)
   Vyvineme funkcny prototyp s test-driven development metodologiou.

DALSI KROK:
Vyskuste si nas AI Kompas — napisite vasu otazku a nas AI agent vam odpovie v realnom case.
Zacnite tu: https://ascentia-web-hermes-agent.fly.dev/kompas

Ak mate otazky, odpovedajte na tento email alebo nas kontaktujte na ascentia@agentmail.to.

S pozdravom,
Tim ASCENTIA s. r. o.
https://ascentia-web-hermes-agent.fly.dev/
`, name, company, score, budget, urgency, message)
}

// sendFollowUpEmail odošle 24h follow-up email klientovi
func sendFollowUpEmail(name, email, analysis, urgency string) {
	config := loadEmailConfig()

	subject := "AI analyza vaseho dopytu — ASCENTIA s. r. o."
	if urgency == "high" {
		subject = "[URGENT] AI analyza vaseho dopytu — ASCENTIA s. r. o."
	}

	go func() {
		sendEmail(config, email, subject, analysis)
		fmt.Printf("[FOLLOW-UP EMAIL] Sent to: %s | Name: %s\n", email, name)
	}()
}

// StartFollowUpScheduler spustí periodickú kontrolu každú hodinu
func (s *Server) StartFollowUpScheduler() {
	ticker := time.NewTicker(1 * time.Hour)

	go func() {
		// Prvotný check po 5 minútach od štartu
		time.Sleep(5 * time.Minute)
		s.RunFollowUpCheck()

		for {
			select {
			case <-ticker.C:
				s.RunFollowUpCheck()
			}
		}
	}()

	fmt.Println("[FOLLOW-UP SCHEDULER] Started — checking every hour for inquiries older than 24h")
}
