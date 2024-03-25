package email

import (
	"github.com/Solidform-labs/newsletter/configs"
	"github.com/Solidform-labs/newsletter/internal/app/newsletter/api/models"
	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/gomail.v2"
)

// SendNewsletter sends a newsletter to all subscribers
func SendNewsletter(recipients []models.Subscriber, subject, body string) error {
	config := configs.GetConfig()
	var d *gomail.Dialer
	if config.Environment == "development" {
		d = gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPassword)
	} else {
		d = &gomail.Dialer{Host: config.SMTPHost, Port: config.SMTPPort}
	}
	s, err := d.Dial()
	defer s.Close()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	for _, r := range recipients {
		m.SetHeader("From", "no-reply@example.com")
		m.SetAddressHeader("To", r.Email, r.Email)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", body)

		if err := gomail.Send(s, m); err != nil {
			log.Errorf("Error sending email to %s: %s", r.Email, err)
			m.Reset()
		}
	}
	log.Info("Newsletter sent")
	return nil
}
