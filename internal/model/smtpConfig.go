package model

import (
	"os"
	"strconv"

	"github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"
)

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	FromName string
}

func NewSMTPConfigFromEnv() (*SMTPConfig, *rest_err.RestErr){
	port := 587
	if v := os.Getenv("ZOHO_SMTP_PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			port = n
		}
	}
	host := os.Getenv("ZOHO_SMTP_HOST")
	if host == "" { host = "smtp.zoho.com" }

	user := os.Getenv("ZOHO_USER")
	pass := os.Getenv("ZOHO_PASS")
	if user == "" || pass == "" {
		return nil, rest_err.NewInternalServerError("missing env: ZOHO_USER or ZOHO_PASS")
	}

	from := os.Getenv("MAIL_FROM_NAME")
	if from == "" { from = "Lavique Dias" }

	return &SMTPConfig{
		Host: host, 
		Port: port, 
		User: user, 
		Pass: pass, 
		FromName: from,
	}, nil
}