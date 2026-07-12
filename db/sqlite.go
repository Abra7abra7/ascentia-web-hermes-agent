package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // Pure Go SQLite driver, no CGO dependency
)

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if not exist
	schemas := []string{
		`CREATE TABLE IF NOT EXISTS chat_messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			session_id TEXT NOT NULL,
			role TEXT NOT NULL,
			message TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS contact_inquiries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			company TEXT DEFAULT '',
			message TEXT NOT NULL,
			voice_path TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS lead_scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			inquiry_id INTEGER,
			score INTEGER DEFAULT 0,
			budget TEXT DEFAULT '',
			urgency TEXT DEFAULT '',
			company_size TEXT DEFAULT '',
			summary TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(inquiry_id) REFERENCES contact_inquiries(id) ON DELETE CASCADE
		);`,
	}

	for _, schema := range schemas {
		_, err := db.Exec(schema)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to execute tables creation schema: %w", err)
		}
	}

	return db, nil
}
