package main

import (
	"ascentia-web/ai"
	"ascentia-web/db"
	"ascentia-web/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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
		log.Println("[Ascentia Engine] Inicializovaný Gemini Provider")
	case "openai", "openrouter":
		aiProvider = &ai.OpenAIProvider{
			APIKey:  os.Getenv("AI_API_KEY"),
			Model:   os.Getenv("AI_MODEL"),
			BaseURL: os.Getenv("AI_BASE_URL"),
		}
		log.Println("[Ascentia Engine] Inicializovaný OpenAI/OpenRouter Provider")
	default:
		aiProvider = &ai.MockProvider{}
		log.Println("[Ascentia Engine] Inicializovaný Mock Provider (Zero Cost)")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "ascentia.db"
	}

	sqliteDB, err := db.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Chyba pri inicializácii SQLite: %v", err)
	}
	defer sqliteDB.Close()

	// Štartovacia migrácia: označ všetky staré dopyty (>24h) ako odoslané follow-up,
	// aby sa neposielali duplicitné emaily po reštarte
	if _, err := sqliteDB.Exec("UPDATE contact_inquiries SET follow_up_sent = 1 WHERE created_at < datetime('now', '-24 hours') AND (follow_up_sent IS NULL OR follow_up_sent = 0)"); err != nil {
		log.Printf("[MIGRACE] Varovanie: nepodarilo sa označiť staré dopyty: %v", err)
	} else {
		log.Println("[MIGRACE] Staré dopyty označené ako follow-up odoslané")
	}

	server, err := handlers.NewServer(sqliteDB, aiProvider)
	if err != nil {
		log.Fatalf("Chyba pri zostavovaní HTTP servera: %v", err)
	}

	server.StartFollowUpScheduler()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})
	mux.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/sitemap.xml")
	})

	mux.HandleFunc("/", server.HandleIndex)
	mux.HandleFunc("/services", server.HandleServices)
	mux.HandleFunc("/process", server.HandleProcess)
	mux.HandleFunc("/privacy", server.HandlePrivacy)
	mux.HandleFunc("/kompas", server.HandleKompas)
	mux.HandleFunc("/voice-inquiry", server.HandleVoice)
	mux.HandleFunc("/faq", server.HandleFAQ)
	mux.HandleFunc("/consultation", server.HandleConsultation)
	mux.HandleFunc("/blog", server.HandleBlog)
	mux.HandleFunc("/pricing", server.HandlePricing)
	mux.HandleFunc("/blog/go-pre-enterprise", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_go_pre_enterprise")
	})
	mux.HandleFunc("/blog/ai-automatizacia-pravnikom", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_ai_pravnikom")
	})
	mux.HandleFunc("/blog/voice-to-crm-case-study", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, r, "blog_voice_crm_case")
	})

	mux.HandleFunc("/admin", server.HandleAdmin)

	mux.HandleFunc("/api/contact", server.HandleContactSubmit)
	mux.HandleFunc("/api/stream", server.HandleStreamingAI)
	mux.HandleFunc("/api/voice-upload", server.HandleVoiceUpload)

	log.Printf("[Ascentia Web] Server úspešne naštartovaný na sluzbe http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Zlyhanie pri štarte web servera: %v", err)
	}
}
