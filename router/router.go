package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mycloud/controller"
)

type Router interface {
	Index() *gin.Engine
}

type router struct {
	uc controller.UserController
	ac controller.AuthController
}

func NewRouter() Router {
	if uc, err := controller.NewUserController(); err != nil {
		log.Fatal("User Controller Assigning is failed")
	} else if ac, err := controller.NewAuthController(); err != nil {
		log.Fatal("Auth Controller Assigning is failed")
	} else {
		return &router{uc, ac}
	}

	return nil

	//return &router{uc}, nil
}

func (r *router) Index() *gin.Engine {
	fmt.Println("router Index func")
	e := gin.Default()

	user := e.Group("/user")
	{
		user.POST("", r.uc.CreateUser)
	}

	auth := e.Group("/auth")
	{
		auth.GET("/callback/google", r.ac.GoogleLogin)
		//auth.POST("", r.Login)
	}

	return e
}
