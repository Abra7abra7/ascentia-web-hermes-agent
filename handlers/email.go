package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// EmailConfig drží SMTP konfiguráciu z environment premenných
type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPass     string
	FromEmail    string
	FromName     string
	NotifyEmail  string
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

// IsEmailConfigured kontroluje či je SMTP správne nakonfigurované
func (c *EmailConfig) IsEmailConfigured() bool {
	return c.SMTPHost != "" && c.SMTPUser != "" && c.SMTPPass != ""
}

// sendEmail odošle email cez Resend REST API (bez potreby SMTP)
func sendEmail(config *EmailConfig, to, subject, body string) error {
	if config.SMTPPass == "" {
		fmt.Printf("[EMAIL SKIP] No API key configured. To: %s | Subject: %s\n", to, subject)
		return nil
	}

	// Resend REST API — from musí byť z verifikovanej domény
	payload := fmt.Sprintf(`{"from":"ASCENTIA Web <ascentia@marianstancik.dev>","to":["%s"],"subject":%q,"text":%q}`,
		to, subject, body)

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+config.SMTPPass)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[EMAIL ERROR] Resend API failed: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[EMAIL SENT] To: %s | Subject: %s | Status: %d\n", to, subject, resp.StatusCode)
	} else {
		fmt.Printf("[EMAIL ERROR] Resend returned %d\n", resp.StatusCode)
	}

	return nil
}

// sendLeadNotification odošle notifikáciu o novom textovom dopyte
func sendLeadNotification(name, email, company, message, source string) {
	config := loadEmailConfig()
	subject := fmt.Sprintf("[ASCENTIA] Nový dopyt od %s (%s)", name, source)
	body := fmt.Sprintf(`Nový dopyt z webu ASCENTIA s. r. o.

Meno: %s
Email: %s
Spoločnosť: %s
Zdroj: %s

Správa:
%s

---
Tento email bol automaticky vygenerovaný systémom Ascentia Web.
AI skórovanie leadu bolo spracované asynchrónne.
`, name, email, company, source, message)

	go func() {
		sendEmail(config, config.NotifyEmail, subject, body)
	}()
}

// sendVoiceLeadNotification odošle notifikáciu o novom hlasovom dopyte
func sendVoiceLeadNotification(name, email, company, voicePath string) {
	config := loadEmailConfig()
	subject := fmt.Sprintf("[ASCENTIA] Nový HLASOVÝ dopyt od %s", name)
	body := fmt.Sprintf(`Nový hlasový dopyt z webu ASCENTIA s. r. o.

Meno: %s
Email: %s
Spoločnosť: %s

Audio súbor: %s

---
Tento email bol automaticky vygenerovaný systémom Ascentia Voice-to-CRM.
AI skórovanie leadu bolo spracované asynchrónne.
`, name, email, company, voicePath)

	go func() {
		sendEmail(config, config.NotifyEmail, subject, body)
	}()

	_ = strings.TrimSpace
}
