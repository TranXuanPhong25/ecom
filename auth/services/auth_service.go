package services

import (
	"fmt"
	"github.com/TranXuanPhong25/ecom/auth/models"
	"golang.org/x/crypto/bcrypt"
)

// var (
//
//	EmailRegex = `/^((?!\.)[\w\-_.]*[^.])(@\w+)(\.\w+(\.\w+)?[^.\W])$/gm`
//
// )
//
//	func IsValidEmailFormat(email string) bool {
//		matched, err := regexp.MatchString(EmailRegex, email)
//		if err != nil {
//			return false
//		}
//		return matched
//	}
func LoginWithEmailAndPassword(email, password string) (models.LoginResponse, error) {
	user, err := GetUserByEmailAndPassword(email, password)
	if err != nil {
		return models.LoginResponse{}, err
	}
	if user == nil {
		return models.LoginResponse{}, fmt.Errorf("user not found")
	}

	token, err := CreateToken(user.UserId)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Token: token,
		User: models.UserInfo{
			UserId: user.UserId,
			Email:  email,
		},
	}, nil
}

func RegisterWithEmailAndPassword(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = CreateUserWithEmailAndPassword(email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
