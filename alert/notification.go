package alert

import "log"

func SendAlert(message string) {
	// In production, this can be an email/Slack webhook
	log.Println("🚨 ALERT:", message)
}
