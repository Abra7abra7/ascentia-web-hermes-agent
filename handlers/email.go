package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// EmailConfig drží SMTP konfiguráciu z environment premenných
type EmailConfig struct {
	SMTPHost    string
	SMTPPort    string
	SMTPUser    string
	SMTPPass    string
	FromEmail   string
	FromName    string
	NotifyEmail string
}

// loadEmailConfig načíta SMTP konfiguráciu z env premenných
func loadEmailConfig() *EmailConfig {
	apiKey := os.Getenv("SMTP_PASS")
	if apiKey == "" {
		// Fallback: hardcoded Resend API key (garantované funkčné aj bez env konfigurácie)
		apiKey = "re_Vo78nycT_26FHEQj4mxP1KtY539fLqDo6"
	}
	return &EmailConfig{
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    apiKey,
		FromEmail:   os.Getenv("FROM_EMAIL"),
		FromName:    "ASCENTIA Web",
		NotifyEmail: "ascentia@agentmail.to",
	}
}

// resendPayload je JSON štruktúra pre Resend API
type resendPayload struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
}

// sendEmail odošle email cez Resend REST API
func sendEmail(config *EmailConfig, to, subject, body string) error {
	if config.SMTPPass == "" {
		fmt.Printf("[EMAIL SKIP] No API key. To: %s | Subject: %s\n", to, subject)
		return nil
	}

	payload := resendPayload{
		From:    "ASCENTIA Web <ascentia@marianstancik.dev>",
		To:      []string{to},
		Subject: subject,
		Text:    body,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("[EMAIL ERROR] JSON marshal failed: %v\n", err)
		return err
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+config.SMTPPass)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[EMAIL ERROR] Resend API request failed: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[EMAIL SENT] To: %s | Subject: %s | Status: %d\n", to, subject, resp.StatusCode)
	} else {
		fmt.Printf("[EMAIL ERROR] Resend returned status %d for to=%s subject=%s\n", resp.StatusCode, to, subject)
	}

	return nil
}

// sendLeadNotification odošle notifikáciu o novom textovom dopyte
func sendLeadNotification(name, email, company, message, source string) {
	config := loadEmailConfig()
	subject := fmt.Sprintf("[ASCENTIA] Novy dopyt od %s (%s)", name, source)
	body := fmt.Sprintf(`Novy dopyt z webu ASCENTIA s. r. o.

Meno: %s
Email: %s
Spolocnost: %s
Zdroj: %s

Sprava:
%s

---
Tento email bol automaticky vygenerovany systemom Ascentia Web.
`, name, email, company, source, message)

	go func() {
		sendEmail(config, config.NotifyEmail, subject, body)
	}()
}

// sendVoiceLeadNotification odošle notifikáciu o novom hlasovom dopyte
func sendVoiceLeadNotification(name, email, company, voicePath string) {
	config := loadEmailConfig()
	subject := fmt.Sprintf("[ASCENTIA] Novy HLASOVY dopyt od %s", name)
	body := fmt.Sprintf(`Novy hlasovy dopyt z webu ASCENTIA s. r. o.

Meno: %s
Email: %s
Spolocnost: %s

Audio subor: %s

---
Tento email bol automaticky vygenerovany systemom Ascentia Voice-to-CRM.
`, name, email, company, voicePath)

	go func() {
		sendEmail(config, config.NotifyEmail, subject, body)
	}()
}
