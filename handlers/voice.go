package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// HandleVoiceUpload prijíma audio nahrávku z prehliadača cez multipart/form-data,
// ukladá ju na disk, vytvára contact_inquiry záznam a asynchrónne kvalifikuje lead cez AI.
func (s *Server) HandleVoiceUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Obmedzenie veľkosti uploadu na 10MB (30s audio v webm/opus je cca 2-4MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	r.ParseMultipartForm(10 << 20)

	name := r.FormValue("name")
	email := r.FormValue("email")
	company := r.FormValue("company")

	if name == "" || email == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Meno a email sú povinné polia"})
		return
	}

	// Prečítanie audio súboru z multipart formulára
	file, header, err := r.FormFile("audio")
	if err != nil {
		// Ak nie je audio, skontrolujeme či prišiel textový prepis
		transcript := r.FormValue("transcript")
		if transcript == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Audio súbor alebo textový prepis je povinný"})
			return
		}

		// Uložíme iba textový dopyt bez audia
		s.saveVoiceInquiry(w, name, email, company, transcript, "")
		return
	}
	defer file.Close()

	// Vytvorenie adresára pre voice nahrávky ak neexistuje
	voiceDir := "/data/voice_uploads"
	if err := os.MkdirAll(voiceDir, 0755); err != nil {
		// Fallback na lokálny adresár
		voiceDir = "voice_uploads"
		os.MkdirAll(voiceDir, 0755)
	}

	// Generovanie unikátneho názvu súboru
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".webm"
	}
	filename := fmt.Sprintf("voice_%d_%s%s", time.Now().Unix(), sanitizeFilename(email), ext)
	voicePath := filepath.Join(voiceDir, filename)

	// Uloženie audio súboru na disk
	dst, err := os.Create(voicePath)
	if err != nil {
		http.Error(w, "Failed to save audio file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to write audio file", http.StatusInternalServerError)
		return
	}

	// Uloženie inquiry do databázy s cestou k audio súboru
	transcript := r.FormValue("transcript")
	s.saveVoiceInquiry(w, name, email, company, transcript, voicePath)
}

// saveVoiceInquiry uloží dopyt do DB a spustí AI kvalifikáciu leadu
func (s *Server) saveVoiceInquiry(w http.ResponseWriter, name, email, company, transcript, voicePath string) {
	if transcript == "" {
		transcript = "Hlasový dopyt od používateľa (audio nahrávka uložená na serveri). Prepis bude spracovaný prostredníctvom AI."
	}

	message := fmt.Sprintf("[Hlasový dopyt] %s", html.EscapeString(transcript))

	res, err := s.DB.Exec("INSERT INTO contact_inquiries (name, email, company, message, voice_path) VALUES (?, ?, ?, ?, ?)",
		name, email, company, message, voicePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Database error"})
		return
	}

	inquiryID, _ := res.LastInsertId()

	// Asynchrónna AI kvalifikácia leadu
	go func(id int64, text string) {
		score, err := s.AI.QualifyLead(context.Background(), text)
		if err == nil {
			s.DB.Exec("INSERT INTO lead_scores (inquiry_id, score, budget, urgency, company_size, summary) VALUES (?, ?, ?, ?, ?, ?)",
				id, score.Score, score.Budget, score.Urgency, score.CompanySize, score.Summary)
		}
	}(inquiryID, fmt.Sprintf("%s. Spoločnosť: %s. Správa: %s", name, company, transcript))

	// Send email notification to ascentia@agentmail.to
	sendVoiceLeadNotification(name, email, company, voicePath)

	// Send confirmation email to client
	sendClientConfirmation(name, email, "Hlasový dopyt (Voice-to-CRM)")

	// Úspešná odpoveď
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Ďakujeme, %s! Váš hlasový dopyt bol úspešne prijatý a spracovaný. Kontaktujeme vás na %s do 24 hodín.", name, email),
	})
}

// sanitizeFilename odstráni nebezpečné znaky z emailu pre bezpečné použitie v názve súboru
func sanitizeFilename(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' || c == '-' {
			result = append(result, c)
		} else {
			result = append(result, '_')
		}
	}
	return string(result)
}
