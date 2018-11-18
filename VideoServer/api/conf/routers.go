package conf

import (
	"github.com/julienschmidt/httprouter"
	"golang/VideoServer/api/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:userName", handlers.Login)
	return router
}
