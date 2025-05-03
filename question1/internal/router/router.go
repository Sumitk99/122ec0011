package router

import (
	"github.com/Sumitk99/122ec0011/question1/internal/handler"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(router *gin.Engine){
	router.POST("numbers/{numberid}", handler.Handler)
}