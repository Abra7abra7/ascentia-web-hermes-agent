package handlers

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
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
	return &EmailConfig{
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		FromEmail:   os.Getenv("FROM_EMAIL"),
		FromName:    "ASCENTIA Web",
		NotifyEmail: "ascentia@agentmail.to",
	}
}

// IsEmailConfigured kontroluje či je SMTP správne nakonfigurované
func (c *EmailConfig) IsEmailConfigured() bool {
	return c.SMTPHost != "" && c.SMTPUser != "" && c.SMTPPass != ""
}

// sendEmail odošle email cez SMTP
func sendEmail(config *EmailConfig, to, subject, body string) error {
	if !config.IsEmailConfigured() {
		// SMTP nie je nakonfigurované — len logneme
		fmt.Printf("[EMAIL SKIP] SMTP not configured. To: %s | Subject: %s\n", to, subject)
		return nil
	}

	from := config.FromEmail
	if from == "" {
		from = config.SMTPUser
	}

	msg := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		config.FromName, from, to, subject, body)

	auth := smtp.PlainAuth("", config.SMTPUser, config.SMTPPass, config.SMTPHost)
	addr := fmt.Sprintf("%s:%s", config.SMTPHost, config.SMTPPort)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Printf("[EMAIL ERROR] Failed to send to %s: %v\n", to, err)
		return err
	}

	fmt.Printf("[EMAIL SENT] To: %s | Subject: %s\n", to, subject)
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
