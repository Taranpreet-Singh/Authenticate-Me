package constants

import "errors"

const (
	User         = "User"
	Success      = "Success"
	Error        = "Error"
	Unauthorized = "Unauthorized"
	LoginSuccess = "Successfully logged in"
	Token        = "token"
)

var (
	ErrUserNotFound      = errors.New("User information not found")
	ErrIncorrectPassword = errors.New("Invalid email or password.")
	ErrJwtToken          = errors.New("Error creating JWT Token")
)
