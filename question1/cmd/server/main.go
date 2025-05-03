package main

import (
	"log"

	"github.com/Sumitk99/122ec0011/question1/internal/config"
	"github.com/Sumitk99/122ec0011/question1/internal/repository"
	routers "github.com/Sumitk99/122ec0011/question1/internal/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
func main(){
	var router *gin.Engine = gin.New() 
	config.LoadEnv()
	ps, err := repository.ConnectToPostgres(os.Getenv("sqlurl"))
	if err != nil {
		log.Fatal(err)
	}

	routers.SetupRoutes(router, err) // func in router package
	router.Run("0.0.0.0:8080") // router from gin package
}