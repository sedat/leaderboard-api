package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sedat/leaderboard-api/handlers"
	"github.com/sedat/leaderboard-api/middleware"
)

func UserRoutes(route *gin.Engine) {

	userRoute := route.Group("/user")
	{
		userRoute.GET("/fill/:amount", middleware.AuthenticationMiddleware(), handlers.FillDB)
		userRoute.POST("/create", handlers.CreateUser)
		userRoute.POST("/login", handlers.LoginUser)
		userRoute.GET("/profile/:id", middleware.AuthenticationMiddleware(), handlers.GetProfile)
		userRoute.POST("/score/submit", middleware.AuthenticationMiddleware(), handlers.SubmitScore)
	}
}
