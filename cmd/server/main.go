package main

import (
	"log"
	"net/http"
	"os"
	"ascentia-web/ai"
	"ascentia-web/db"
	"ascentia-web/handlers"
)

func main() {
	// Načítanie portu z prostredia alebo rozumný default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Výber konfigurácie AI Providera zo špecifikácie .env
	providerType := os.Getenv("AI_PROVIDER")
	if providerType == "" {
		providerType = "mock" // Safe default pre lokálne testovanie za 0€
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

	// Inicializácia DB
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "ascentia.db"
	}

	sqliteDB, err := db.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Chyba pri inicializácii SQLite: %v", err)
	}
	defer sqliteDB.Close()

	// Vytvorenie dedikovaného servera s handlermi a šablónami
	server, err := handlers.NewServer(sqliteDB, aiProvider)
	if err != nil {
		log.Fatalf("Chyba pri zostavovaní HTTP servera: %v", err)
	}

	// Štart automatizovaného follow-up schedulera (kontroluje dopyty staršie ako 24h)
	server.StartFollowUpScheduler()

	// Definícia prísne optimalizovaných rout
	mux := http.NewServeMux()

	// Statické súbory
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Presmerovanie /robots.txt a /sitemap.xml na statické súbory (SEO/GEO)
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})
	mux.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/sitemap.xml")
	})

	// Trasy pre jednotlivé stránky
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

	// Admin dashboard
	mux.HandleFunc("/admin", server.HandleAdmin)

	// POST trasy a zber leadov
	mux.HandleFunc("/api/contact", server.HandleContactSubmit)
	mux.HandleFunc("/api/stream", server.HandleStreamingAI)
	mux.HandleFunc("/api/voice-upload", server.HandleVoiceUpload)

	log.Printf("[Ascentia Web] Server úspešne naštartovaný na sluzbe http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Zlyhanie pri štarte web servera: %v", err)
	}
}
