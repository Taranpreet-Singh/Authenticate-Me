package service

import (
	"authenticate-me/src/initializers"
	"authenticate-me/src/models"
	"authenticate-me/src/utils"
	"authenticate-me/src/utils/constants"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(body *models.UserRequestBody) (*models.User, *models.CustomError) {
	//* Generate hash for the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return nil, &models.CustomError{ErrorType: constants.Error, Error: err}
	}

	//* Register user in DB
	user := &models.User{ID: uuid.New(), Email: body.Email, Password: string(hash)}
	userObj := initializers.DB.Create(&user)
	if userObj.Error != nil {
		return nil, &models.CustomError{ErrorType: constants.Error, Error: userObj.Error}
	}
	return user, nil
}

func LoginUser(body *models.UserRequestBody) (string, *models.CustomError) {
	jwtService := &jwtServices{
		privateKey: utils.GetSecretKey(),
	}
	//* Get the requested user
	user := new(models.User)
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		return "", &models.CustomError{ErrorType: constants.Error, Error: constants.ErrUserNotFound}
	}

	//* Compare with recieved request body
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return "", &models.CustomError{ErrorType: constants.Unauthorized, Error: constants.ErrIncorrectPassword}
	}

	token, err := jwtService.GetJwtToken(user)
	if err != nil {
		return "", &models.CustomError{ErrorType: constants.Error, Error: constants.ErrJwtToken}
	}
	return token, nil
}
