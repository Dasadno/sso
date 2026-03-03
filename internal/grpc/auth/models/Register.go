package models

// register request model
type Register struct {
	Email             string `validate:"required,email,min=5,max=50"`
	Password          string `validate:"required,email,min=8,max=50,passwd"`
	PasswordConfirmed string `validate:"required,eqfield=Password,email,min=8,max=50,passwd"`
}
