package service

import (
	"github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-sendEmail-go/internal/model"
)

func NewSendEmailService(sender model.SMTPSender) SendEmailSerivice {
	return &sendEmailService{
		sender: sender,
	}
}

type sendEmailService struct {
	sender model.SMTPSender
}

type SendEmailSerivice interface {
	SendEmail(emailRequest model.EmailRequest) (string, *rest_err.RestErr)
}