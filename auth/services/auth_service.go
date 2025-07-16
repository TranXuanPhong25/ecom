package services

import (
	"github.com/TranXuanPhong25/ecom/auth/models"
)

func LoginWithEmailAndPassword(email, password string) (models.LoginResponse, error) {
	user, err := GetUserByEmailAndPassword(email, password)
	if err != nil {
		return models.LoginResponse{}, err
	}

	token, err := CreateToken(user.UserId)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Token: token,
		User: models.UserInfo{
			UserId: user.UserId,
			Email:  user.Email,
		},
	}, nil
}

func RegisterWithEmailAndPassword(email, password string) error {

	err := CreateUserWithEmailAndPassword(email, password)
	if err != nil {
		return err
	}

	return nil
}

func GetCurrentUser(userId string) (*models.UserInfo, error) {
	user, err := GetUserById(userId)
	if err != nil {
		return &models.UserInfo{}, err
	}

	return user, nil
}
