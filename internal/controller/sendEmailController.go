package controller

import (
	"github.com/LaviqueDias/api-sendEmail-go/internal/service"
	"github.com/gin-gonic/gin"
)

func NewSendEmailController(service service.SendEmailSerivice) SendEmailController {
	return &sendEmailController{
		service: service,
	}
}

type sendEmailController struct {
	service service.SendEmailSerivice
}

type SendEmailController interface {
	SendEmail(c *gin.Context)
}