package routes

import (
	"e-learning-platform/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTMiddleware())

	return router
}

func StartService(route *gin.Engine) {
	route.Run()
}
