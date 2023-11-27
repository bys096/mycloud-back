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
}

func NewRouter() Router {
	if uc, err := controller.NewUserController(); err != nil {
		log.Fatal("User Controller Assigning is failed")
	} else {
		return &router{uc}
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

	//auth := e.Group("/auth")
	//{
	//	user.POST("", r.Login)
	//}

	return e
}
