package db

import (
	"database/sql"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	tempDB := "test_ascentia.db"
	defer os.Remove(tempDB)

	db, err := InitDB(tempDB)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()

	// Verify tables
	tables := []string{"chat_messages", "contact_inquiries", "lead_scores"}
	for _, table := range tables {
		var name string
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?;", table).Scan(&name)
		if err == sql.ErrNoRows {
			t.Errorf("Table %s was not created", table)
		} else if err != nil {
			t.Fatalf("Failed to query sqlite_master: %v", err)
		}
	}
}
