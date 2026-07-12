package handlers

import (
	"context"
	"fmt"
	"net/http"
)

func (s *Server) HandleContactSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	company := r.FormValue("company")
	message := r.FormValue("message")

	if name == "" || email == "" || message == "" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<div class="error-msg">Prosím, vyplňte všetky povinné polia (Meno, Email a Správu).</div>`))
		return
	}

	// Save inquiry
	res, err := s.DB.Exec("INSERT INTO contact_inquiries (name, email, company, message) VALUES (?, ?, ?, ?)",
		name, email, company, message)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	inquiryID, _ := res.LastInsertId()

	// Qualify Lead via AI asynchronously + send email notification
	go func(id int64, text string) {
		score, err := s.AI.QualifyLead(context.Background(), text)
		if err == nil {
			_, _ = s.DB.Exec("INSERT INTO lead_scores (inquiry_id, score, budget, urgency, company_size, summary) VALUES (?, ?, ?, ?, ?, ?)",
				id, score.Score, score.Budget, score.Urgency, score.CompanySize, score.Summary)
		}
	}(inquiryID, fmt.Sprintf("%s. Spoločnosť: %s. Správa: %s", name, company, message))

	// Send email notification to ascentia@agentmail.to
	sendLeadNotification(name, email, company, message, "B2B formulár")

	// Send confirmation email to client
	sendClientConfirmation(name, email, "B2B formulár")

	// Pre HTMX vrátime pekný úspešný fragment, inak zobrazenie správy
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`
		<div class="success-box">
			<h3>✓ Dopyt bol úspešne odoslaný!</h3>
			<p>Ďakujeme, <strong>%s</strong>. Vašu žiadosť sme zaznamenali a naši inžinieri z Ascentia s. r. o. vás budú kontaktovať emailom na <strong>%s</strong> do 24 hodín.</p>
		</div>
	`, name, email)))
}
