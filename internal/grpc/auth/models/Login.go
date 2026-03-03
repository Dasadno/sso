package models

// login request model
type Login struct {
	Email    string `validate:"required,min=5,max=50,email"`
	Password string `validate:"required,min=8,max=64,passwd"`
}
