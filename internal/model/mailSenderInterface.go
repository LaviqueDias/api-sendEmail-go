package model

import "github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"

type MailSender interface {
	Send(emailRequest EmailRequest) *rest_err.RestErr
}