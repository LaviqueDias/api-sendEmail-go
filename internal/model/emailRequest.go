package model

type EmailRequest struct {
	Name        string `json:"name"        binding:"required,notblank,mintrim=2"`
	Email       string `json:"email"       binding:"required,email"`
	Message     string `json:"message"     binding:"required,notblank,mintrim=5"`
}

