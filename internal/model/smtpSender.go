package model 

import (
	"fmt"
	"github.com/go-mail/mail/v2"
	"time"

	"github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"
)



type SMTPSender struct {
	config SMTPConfig
}

func NewSMTPSender(config *SMTPConfig) *SMTPSender {
	return &SMTPSender{
		config: *config,
	}
}

func (s *SMTPSender) Send(emailRequest EmailRequest) (string, *rest_err.RestErr) {
    replyTo := ""                 

    m := mail.NewMessage()
    m.SetHeader("From", s.config.User)
    m.SetHeader("To", emailRequest.Email)
    if replyTo != "" { m.SetHeader("Reply-To", replyTo) }
    m.SetHeader("Subject", fmt.Sprintf("Hello %s", emailRequest.Name))

    html := fmt.Sprintf(`
        <div style="font-family:Arial,Helvetica,sans-serif;font-size:14px;color:#222">
            <h2>Email sent via API</h2>
            <p><b>Nome:</b> %s</p>
            <p><b>Email:</b> %s</p>
            <p><b>Mensagem:</b></p>
            <p style="white-space:pre-wrap">%s</p>
            <br>
            <h4>API repo: github.com/LaviqueDias/api-sendEmail-go</h4>
        </div>`,
        emailRequest.Name, emailRequest.Email, emailRequest.Message,
    )
    m.SetBody("text/html", html)

    d := mail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Pass)
    d.StartTLSPolicy = mail.MandatoryStartTLS
    d.Timeout = 10 * time.Second

    if err := d.DialAndSend(m); err != nil {
        fmt.Printf("[EMAIL] smtp send failed: %v\n", err)
        return "", rest_err.NewInternalServerError("smtp send failed: " + err.Error())
    }
    return html, nil
}

