package handlers

import (
	"database/sql"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

// setupFollowUpTestDB vytvorí izolovanú SQLite DB s potrebnými tabuľkami
// a vráti *sql.DB plus cleanup funkciu.
func setupFollowUpTestDB(t *testing.T) (*sql.DB, func()) {
	t.Helper()
	tempDB := "test_followup.db"
	// cleanup starého súboru
	_ = os.Remove(tempDB)

	db, err := sql.Open("sqlite", tempDB)
	if err != nil {
		t.Fatalf("Failed to open test db: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping test db: %v", err)
	}

	schemas := []string{
		`CREATE TABLE contact_inquiries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			company TEXT DEFAULT '',
			message TEXT NOT NULL,
			voice_path TEXT DEFAULT '',
			follow_up_sent INTEGER DEFAULT 0,
			sent_at DATETIME DEFAULT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE lead_scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			inquiry_id INTEGER,
			score INTEGER DEFAULT 0,
			budget TEXT DEFAULT '',
			urgency TEXT DEFAULT '',
			company_size TEXT DEFAULT '',
			summary TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}
	for _, s := range schemas {
		if _, err := db.Exec(s); err != nil {
			t.Fatalf("Failed to create schema: %v", err)
		}
	}

	cleanup := func() {
		db.Close()
		os.Remove(tempDB)
	}
	return db, cleanup
}

// insertInquiry vloží dopyt so zadaným emailom a follow_up_sent príznakom.
func insertInquiry(t *testing.T, db *sql.DB, email string, followUpSent int) int64 {
	t.Helper()
	res, err := db.Exec(
		`INSERT INTO contact_inquiries (name, email, company, message, follow_up_sent, created_at)
		 VALUES (?, ?, 'TestCo', 'test message', ?, datetime('now', '-48 hours'))`,
		"Test User", email, followUpSent,
	)
	if err != nil {
		t.Fatalf("Failed to insert inquiry: %v", err)
	}
	id, _ := res.LastInsertId()
	return id
}

// seedLeadScore vloží lead_score záznam pre daný inquiry.
func seedLeadScore(t *testing.T, db *sql.DB, inquiryID int64, score int64) {
	t.Helper()
	_, err := db.Exec(
		`INSERT INTO lead_scores (inquiry_id, score, budget, urgency)
		 VALUES (?, ?, 'High-Ticket Enterprise AI', 'high')`,
		inquiryID, score,
	)
	if err != nil {
		t.Fatalf("Failed to seed lead score: %v", err)
	}
}

// countPendingFollowUps vráti počet dopytov s follow_up_sent = 0.
func countPendingFollowUps(t *testing.T, db *sql.DB) int {
	t.Helper()
	var n int
	if err := db.QueryRow("SELECT COUNT(*) FROM contact_inquiries WHERE follow_up_sent = 0").Scan(&n); err != nil {
		t.Fatalf("Failed to count pending: %v", err)
	}
	return n
}

// TestRunFollowUpCheck_NoDoubleSend overí, že jeden dopyt sa odošle presne raz
// aj pri viacnásobnom volaní RunFollowUpCheck (idempotencia cez atomický UPDATE...RETURNING).
func TestRunFollowUpCheck_NoDoubleSend(t *testing.T) {
	db, cleanup := setupFollowUpTestDB(t)
	defer cleanup()

	_ = insertInquiry(t, db, "client@example.com", 0)
	// seed lead score aby analýza mala dáta
	// (RunFollowUpCheck robí LEFT JOIN, bez score funguje s COALESCE)

	srv := &Server{DB: db}

	// Prvé spustenie — malo odoslať 1 a označiť ako odoslané
	srv.RunFollowUpCheckWithSender(func(name, email, analysis, urgency string) {
		// simulácia odoslania — nič nerobíme
	})
	if got := countPendingFollowUps(t, db); got != 0 {
		t.Errorf("After first run, expected 0 pending, got %d", got)
	}

	// Druhé spustenie — NESMIE odoslať znova (žiadna duplicita)
	var sendCalls int
	srv.RunFollowUpCheckWithSender(func(name, email, analysis, urgency string) {
		sendCalls++
	})
	if sendCalls != 0 {
		t.Errorf("Second run sent %d emails, expected 0 (no duplicates)", sendCalls)
	}
	if got := countPendingFollowUps(t, db); got != 0 {
		t.Errorf("After second run, expected 0 pending, got %d", got)
	}
}

// TestRunFollowUpCheck_OnlyOldInquiries overí, že nové dopyty (<24h) sa neposielajú.
func TestRunFollowUpCheck_OnlyOldInquiries(t *testing.T) {
	db, cleanup := setupFollowUpTestDB(t)
	defer cleanup()

	// Starý dopyt (>24h) — mal sa poslať
	oldID := insertInquiry(t, db, "old@example.com", 0)
	seedLeadScore(t, db, oldID, 85)
	// Nový dopyt (<24h) — NESMIE sa poslať
	_, err := db.Exec(
		`INSERT INTO contact_inquiries (name, email, company, message, follow_up_sent, created_at)
		 VALUES ('New User', 'new@example.com', 'TestCo', 'new msg', 0, datetime('now'))`,
	)
	if err != nil {
		t.Fatalf("Failed to insert new inquiry: %v", err)
	}

	srv := &Server{DB: db}
	var sentEmails []string
	srv.RunFollowUpCheckWithSender(func(name, email, analysis, urgency string) {
		sentEmails = append(sentEmails, email)
	})

	if len(sentEmails) != 1 {
		t.Fatalf("Expected exactly 1 email (old inquiry), got %d: %v", len(sentEmails), sentEmails)
	}
	if sentEmails[0] != "old@example.com" {
		t.Errorf("Expected email to old@example.com, got %s", sentEmails[0])
	}
}

// TestRunFollowUpCheck_MultipleRecipientsOnceEach overí, že 3 dopyty na ten istý
// email sa odošlú presne 3× (raz každý), nie 3× každý pri jednom behu.
func TestRunFollowUpCheck_MultipleRecipientsOnceEach(t *testing.T) {
	db, cleanup := setupFollowUpTestDB(t)
	defer cleanup()

	// 3 samostatné staré dopyty na ten istý email (scenario: 3 záznamy v DB)
	for i := 0; i < 3; i++ {
		id := insertInquiry(t, db, "stancikmarian8@gmail.com", 0)
		seedLeadScore(t, db, id, 80+int64(i))
	}

	srv := &Server{DB: db}
	var sentEmails []string
	srv.RunFollowUpCheckWithSender(func(name, email, analysis, urgency string) {
		sentEmails = append(sentEmails, email)
	})

	if len(sentEmails) != 3 {
		t.Errorf("Expected exactly 3 emails (one per inquiry), got %d: %v", len(sentEmails), sentEmails)
	}
	// Žiadny príjemca nesmie dostať viac než 1 v jednom behu
	counts := map[string]int{}
	for _, e := range sentEmails {
		counts[e]++
	}
	if counts["stancikmarian8@gmail.com"] != 3 {
		t.Errorf("Expected 3 distinct sends to stancikmarian8@gmail.com (one per inquiry), got %d", counts["stancikmarian8@gmail.com"])
	}

	// Druhý beh — žiadne nové odoslanie
	sentEmails = nil
	srv.RunFollowUpCheckWithSender(func(name, email, analysis, urgency string) {
		sentEmails = append(sentEmails, email)
	})
	if len(sentEmails) != 0 {
		t.Errorf("Second run sent %d emails, expected 0", len(sentEmails))
	}
}
