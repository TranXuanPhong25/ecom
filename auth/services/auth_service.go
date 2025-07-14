package services

import (
	"github.com/TranXuanPhong25/ecom/auth/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func IsValidEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func LoginWithEmailAndPassword(email, password string) (models.LoginResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.LoginResponse{}, err
	}
	userInfo, err := createUserWithEmailAndPassword(email, string(hashedPassword))
	if err != nil {
		return models.LoginResponse{}, err
	}

	token, err := createToken(userInfo.UserId)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return models.LoginResponse{
		Token: token,
		User:  userInfo,
	}, nil
}

func RegisterWithEmailAndPassword(c echo.Context) error {
	return nil
}
