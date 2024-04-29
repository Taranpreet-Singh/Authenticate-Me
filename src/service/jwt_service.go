package service

import (
	"authenticate-me/src/initializers"
	"authenticate-me/src/models"
	"authenticate-me/src/utils/constants"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TODO: Implementation pending
type JWTService interface {
	TokenGenerate(user *models.User) string
	TokenValidate(token string) (*jwt.Token, models.CustomError)
	GetJwtToken(user *models.User)
}

type jwtServices struct {
	privateKey []byte
}

type authCustomClaims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"name"`
	jwt.StandardClaims
}

func (service *jwtServices) TokenGenerate(user *models.User) string {
	claims := &authCustomClaims{
		user.ID,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//* encoded the web token
	tokenString, err := token.SignedString(service.privateKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (service *jwtServices) GetJwtToken(user *models.User) (string, *models.CustomError) {
	//* Create JWT Token authentication and return it as a request
	token := service.TokenGenerate(user)
	jwt := &models.JwtToken{ID: uuid.New(), UserID: user.ID, Token: token}
	jwtObj := initializers.DB.Create(&jwt)
	if jwtObj.Error != nil {
		return "", &models.CustomError{ErrorType: constants.Error, Error: constants.ErrJwtToken}
	}

	//* Clear previous logins that are more than login limit
	clearOldLogins(user.ID)
	return token, nil
}

func clearOldLogins(userId uuid.UUID) {
	loginLimit, err := strconv.Atoi(os.Getenv("LOGIN_LIMIT"))
	if err != nil {
		panic(err)
	}
	var totalLogins int64
	initializers.DB.Model(&models.JwtToken{}).Where("user_id = ?", userId).Count(&totalLogins)
	if int(totalLogins) > loginLimit {
		// TODO: add a cronjob to delete the entries after 30days
		initializers.DB.Exec(constants.DeleteOldTokens)
	}
}
