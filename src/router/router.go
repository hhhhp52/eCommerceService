package router

import (
	"eCommerceService/src/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", services.Register)
	router.POST("/login", services.Login)
	router.POST("/verify-email", services.VerifyEmail)
	router.GET("/recommendation", services.GetRecommendation)

	return router
}
