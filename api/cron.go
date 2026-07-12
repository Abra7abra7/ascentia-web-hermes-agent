package api

import (
	"fmt"
	"net/http"
	"os"
	"ascentia-web/ai"
	"ascentia-web/db"
	"ascentia-web/handlers"
)

// CronHandler je Vercel Cron endpoint pre follow-up scheduler
func CronHandler(w http.ResponseWriter, r *http.Request) {
	// Overenie, že volanie prichádza z Vercel Cron
	authHeader := r.Header.Get("Authorization")
	if authHeader != "Bearer "+os.Getenv("CRON_SECRET") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	providerType := os.Getenv("AI_PROVIDER")
	if providerType == "" {
		providerType = "mock"
	}

	var aiProvider ai.Provider
	switch providerType {
	case "gemini":
		aiProvider = &ai.GeminiProvider{
			APIKey:  os.Getenv("AI_API_KEY"),
			Model:   os.Getenv("AI_MODEL"),
			BaseURL: os.Getenv("AI_BASE_URL"),
		}
	case "openai", "openrouter":
		aiProvider = &ai.OpenAIProvider{
			APIKey:  os.Getenv("AI_API_KEY"),
			Model:   os.Getenv("AI_MODEL"),
			BaseURL: os.Getenv("AI_BASE_URL"),
		}
	default:
		aiProvider = &ai.MockProvider{}
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "/tmp/ascentia.db"
	}

	sqliteDB, err := db.InitDB(dbPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB error: %v", err), http.StatusInternalServerError)
		return
	}
	defer sqliteDB.Close()

	server, err := handlers.NewServer(sqliteDB, aiProvider)
	if err != nil {
		http.Error(w, fmt.Sprintf("Server error: %v", err), http.StatusInternalServerError)
		return
	}

	server.RunFollowUpCheck()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","message":"follow-up check completed"}`))
}
