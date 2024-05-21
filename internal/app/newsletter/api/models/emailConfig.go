package models

type EmailConfig struct {
	Subject string   `json:subject`
	Body    string   `json:body`
	Ids     []string `json:ids`
}
