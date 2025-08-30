package service

import (
	"github.com/LaviqueDias/api-sendEmail-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-sendEmail-go/internal/model"
)

func (ss *sendEmailService) SendEmail(emailRequest model.EmailRequest) (string, *rest_err.RestErr) {
	html, err := ss.sender.Send(emailRequest)
	if err != nil {
		return "", err
	}

	return html, nil
}