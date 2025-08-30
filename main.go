package main

import (
	"log"

	"github.com/LaviqueDias/api-sendEmail-go/internal/controller"
	"github.com/LaviqueDias/api-sendEmail-go/internal/model"
	"github.com/LaviqueDias/api-sendEmail-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	smtpConfig, err := model.NewSMTPConfigFromEnv()
	if err != err {
		log.Fatal(err)
	}
	
	smtpSender := model.NewSMTPSender(smtpConfig)
	sendEmailService := service.NewSendEmailService(*smtpSender)
	sendEmailController := controller.NewSendEmailController(sendEmailService)
	
	router := gin.Default()
	api := router.Group("/api")
	
	emailGroup := api.Group("/email")

	emailGroup.POST("/send", sendEmailController.SendEmail)

	if err := router.Run(":8080"); err != nil {
		log.Fatal()
	}
}