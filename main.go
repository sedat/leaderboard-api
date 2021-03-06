package main

import (
	"fmt"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sedat/leaderboard-api/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env couldn't open")
	}
}

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(helmet.Default())

	routes.LeaderBoardRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
