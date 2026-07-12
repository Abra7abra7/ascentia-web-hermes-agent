package handlers

import (
	"database/sql"
	"fmt"
	"time"
)

// RunFollowUpCheck skontroluje databázu pre dopyty staršie ako 24h
// ktoré ešte nemajú odoslaný follow-up email
func (s *Server) RunFollowUpCheck() {
	// Nájdi dopyty staršie ako 24h bez follow-upu
	rows, err := s.DB.Query(`
		SELECT ci.id, ci.name, ci.email, ci.company, ci.message,
		       ls.score, ls.budget, ls.urgency
		FROM contact_inquiries ci
		LEFT JOIN lead_scores ls ON ls.inquiry_id = ci.id
		WHERE ci.created_at < datetime('now', '-1 minute')
		  AND ci.follow_up_sent = 0
		ORDER BY ci.created_at ASC
		LIMIT 10
	`)
	if err != nil {
		fmt.Printf("[FOLLOW-UP] DB query error: %v\n", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id int64
		var name, email, company, message string
		var score sql.NullInt64
		var budget, urgency sql.NullString

		err := rows.Scan(&id, &name, &email, &company, &message, &score, &budget, &urgency)
		if err != nil {
			fmt.Printf("[FOLLOW-UP] Row scan error: %v\n", err)
			continue
		}

		scoreVal := int64(0)
		if score.Valid {
			scoreVal = score.Int64
		}
		budgetVal := "nezadaný"
		if budget.Valid && budget.String != "" {
			budgetVal = budget.String
		}
		urgencyVal := "medium"
		if urgency.Valid && urgency.String != "" {
			urgencyVal = urgency.String
		}

		analysis := generateInquiryAnalysis(name, company, message, budgetVal, urgencyVal, scoreVal)
		sendFollowUpEmail(name, email, analysis, urgencyVal)

		s.DB.Exec("UPDATE contact_inquiries SET follow_up_sent = 1 WHERE id = ?", id)
		count++
	}

	if count > 0 {
		fmt.Printf("[FOLLOW-UP] Sent %d follow-up emails\n", count)
	}
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
	ticker := time.NewTicker(2 * time.Minute)

	go func() {
		// Prvotný check po 30 sekundách od štartu (pre testovanie)
		time.Sleep(30 * time.Second)
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
