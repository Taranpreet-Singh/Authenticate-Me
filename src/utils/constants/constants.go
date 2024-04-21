package constants

import "errors"

const (
	User         = "User"
	Success      = "Success"
	Error        = "Error"
	Unauthorized = "Unauthorized"
	LoginSuccess = "Successfully logged in"
)

var (
	ErrUserNotFound      = errors.New("User information not found")
	ErrIncorrectPassword = errors.New("Invalid email or password.")
)
