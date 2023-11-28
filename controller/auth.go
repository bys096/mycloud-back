package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mycloud/config"
	"mycloud/model/repository"
	"net/http"
)

type AuthController interface {
	GoogleLogin(c *gin.Context)
}

type authController struct {
	um *repository.User
}

func NewAuthController() (AuthController, error) {
	um := repository.NewUserModel()
	return &authController{um}, nil
}

func (ac *authController) GoogleLogin(c *gin.Context) {
	log.Println("Google Login Function is called")
	//code := c.Param("code")
	code := c.Query("code")
	google := config.NewGoogle()
	t, _ := google.GenerateToken(code)
	c.JSON(http.StatusOK, t)
}
