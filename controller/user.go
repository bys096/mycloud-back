package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mycloud/errors"
	"mycloud/model/repository"
	"net/http"
	"time"
)

type UserController interface {
	CreateUser(c *gin.Context)
}

type userController struct {
	um *repository.User
}

func NewUserController() (UserController, error) {
	return &userController{}, nil
}

func (uc *userController) CreateUser(c *gin.Context) {
	user := repository.NewUserModel()
	err := c.ShouldBind(&user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	log.Println(user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	_, err = uc.um.Save(user)
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": errors.GetCode(err),
		"msg":  err.Error(),
	})

}
