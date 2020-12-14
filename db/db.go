package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client          *mongo.Client
	UsersCollection *mongo.Collection
}

func (mongoDb *MongoDB) ConnectDB(ctx context.Context) {
	defer ctx.Done()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		log.Fatal(err)
	}

	mongoDb.Client = client
	err = client.Database("leaderboard_db").CreateCollection(ctx, "users")
	mongoDb.UsersCollection = client.Database("leaderboard_db").Collection("users")

}
