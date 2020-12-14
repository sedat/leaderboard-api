package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sedat/leaderboard-api/db"
	"github.com/sedat/leaderboard-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context
var redisClient *redis.Client
var usersCollection *mongo.Collection
var leaderboardCollection *mongo.Collection
var lastRank int

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env couldn't open")
	}
	var mongoDb db.MongoDB
	ctx = context.Background()
	// opt, err := redis.ParseURL(os.Getenv("REDIS"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS"),
	})
	mongoDb.ConnectDB(ctx)

	usersCollection = mongoDb.UsersCollection

	fmt.Println("Initializing redis with leaderboard, please wait. Depending on the size of mongo collection. This will take time.")
	services.InitRedis(ctx, usersCollection, redisClient)
	fmt.Println("Leaderboard initialized.")

}
