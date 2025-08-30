package controller

import (
	"log"
	"net/http"

	validation "github.com/LaviqueDias/api-sendEmail-go/internal/configuration/valdiation"
	"github.com/LaviqueDias/api-sendEmail-go/internal/model"
	"github.com/gin-gonic/gin"
)

func (cs *sendEmailController) SendEmail(c *gin.Context) {
	var emailRequest model.EmailRequest

	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		errRest := validation.ValidateRequestError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	
	html, err := cs.service.SendEmail(emailRequest)
	if err != nil {
		log.Printf("[EMAIL] send failed: %v\n", err.Error())
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email enviado com sucesso",
		"emailBody": html,
	})
}