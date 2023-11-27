package controller

import "mycloud/model/repository"

type AuthController interface {
	//Login()
}

type authController struct {
	um *repository.User
}

func NewAuthController() (AuthController, error) {
	um := repository.NewUserModel()
	return &authController{um}, nil
}
