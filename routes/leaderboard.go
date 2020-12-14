package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sedat/leaderboard-api/handlers"
	"github.com/sedat/leaderboard-api/middleware"
)

func LeaderBoardRoutes(route *gin.Engine) {
	leaderboardRoute := route.Group("/leaderboard")
	leaderboardRoute.Use(middleware.AuthenticationMiddleware())
	{
		leaderboardRoute.GET("/", handlers.GetLeaderboard)
		leaderboardRoute.GET("/limit/:limit", handlers.GetLeaderboardWithLimit)
		leaderboardRoute.GET("/country/:country/", handlers.GetLeaderboardByCountry)
		leaderboardRoute.GET("/country/:country/limit/:limit", handlers.GetLeaderboardByCountryWithLimit)
		leaderboardRoute.GET("/range/:start/:end", handlers.GetLeaderboardInRange)
	}
}
