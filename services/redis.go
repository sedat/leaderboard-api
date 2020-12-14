package services

import (
	"context"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/sedat/leaderboard-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRedis(ctx context.Context, usersCollection *mongo.Collection, redisClient *redis.Client) {
	pipeline := []bson.M{
		{
			"$sort": bson.M{
				"points": -1,
			},
		},
	}

	redisClient.FlushAll(ctx)

	cursor, err := usersCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	cursor.All(ctx, &users)

	for _, v := range users {
		_, err := redisClient.ZAdd(ctx, "leaderboard", &redis.Z{
			Score:  float64(v.Points),
			Member: v.ID.Hex(),
		}).Result()
		if err != nil {
			log.Fatal(err)
		}
		_, err = redisClient.ZAdd(ctx, "leaderboard:"+strings.ToLower(v.Country), &redis.Z{
			Score:  float64(v.Points),
			Member: v.ID.Hex(),
		}).Result()

		if err != nil {
			log.Fatal(err)
		}
	}
}
