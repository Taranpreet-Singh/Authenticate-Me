package services

import (
	"authenticate-me/src/initializers"
	"authenticate-me/src/models"
	"authenticate-me/src/utils/constants"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(body *models.UserRequestBody) (*models.User, *models.CustomError) {
	// Generate hash for the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, &models.CustomError{ErrorType: constants.Error, Error: err}
	}

	// Register user in DB
	user := &models.User{ID: uuid.New(), Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return nil, &models.CustomError{ErrorType: constants.Error, Error: result.Error}
	}
	return user, nil
}

func LoginUser(body *models.UserRequestBody) (*models.User, *models.CustomError) {
	// Get the requested user
	user := new(models.User)
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.New() {
		return nil, &models.CustomError{ErrorType: constants.Error, Error: constants.ErrUserNotFound}
	}

	// Compare with recieved request body
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, &models.CustomError{ErrorType: constants.Unauthorized, Error: constants.ErrIncorrectPassword}
	}

	return user, nil
}
