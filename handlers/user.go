package handlers

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sedat/leaderboard-api/models"
	"github.com/sedat/leaderboard-api/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func FillDB(c *gin.Context) {
	amount := c.Param("amount")
	err := usersCollection.Drop(ctx)
	if err != nil {
		c.JSON(200, gin.H{"msg": err.Error()})
		fmt.Println("err1")
		return
	}
	_, err = usersCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "display_name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		c.JSON(200, gin.H{"msg": err.Error()})
		return
	}

	cmd := exec.Command("mgodatagen", "-f", "datagen"+amount+"k.json")

	err = cmd.Run()
	if err != nil {
		c.JSON(200, gin.H{"msg": err.Error()})
		return
	}

	services.InitRedis(ctx, usersCollection, redisClient)

	c.JSON(200, gin.H{"msg": "Database created and redis is initialized with leaderboard"})

}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	user.ID = primitive.NewObjectID()
	user.Points = 0
	user.Country = strings.ToUpper(user.Country)
	if err != nil {
		c.JSON(400, gin.H{"msg": "bad request body"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	_, err = usersCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	redisClient.ZAdd(ctx, "leaderboard", &redis.Z{
		Score:  float64(user.Points),
		Member: user.ID.Hex(),
	})

	redisClient.ZAdd(ctx, "leaderboard:"+strings.ToLower(user.Country), &redis.Z{
		Score:  float64(user.Points),
		Member: user.ID.Hex(),
	})

	rank, err := redisClient.ZRevRank(ctx, "leaderboard", user.ID.Hex()).Result()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	user.Rank = int(rank + 1)
	c.JSON(201, user)
}

func LoginUser(c *gin.Context) {
	var userLogin services.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"msg": "user credentials are malformed"})
		return
	}
	var user models.User
	err := usersCollection.FindOne(ctx, bson.D{{"display_name", userLogin.DisplayName}}).Decode(&user)
	if err != nil {
		c.JSON(404, gin.H{"msg": err.Error()})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
		return
	}
	userID := user.ID.Hex()
	userLogin.UserID = userID

	token, err := userLogin.GenerateToken()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func GetProfile(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": "malformed id"})
		return
	}
	var user models.User
	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{"password": 0})
	if err := usersCollection.FindOne(ctx, bson.M{"_id": id}, findOptions).Decode(&user); err != nil {
		c.JSON(404, gin.H{"msg": "user not found"})
		return
	}
	rank, err := redisClient.ZRevRank(ctx, "leaderboard", user.ID.Hex()).Result()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	user.Rank = int(rank + 1)

	c.JSON(200, user)
}

func SubmitScore(c *gin.Context) {
	var userScore models.UserScore
	err := c.BindJSON(&userScore)
	if err != nil {
		c.JSON(404, gin.H{"msg": err.Error()})
		return
	}
	var user models.User
	err = usersCollection.FindOneAndUpdate(ctx, bson.M{"_id": userScore.ID}, bson.D{{"$inc", bson.D{{"points", userScore.Score}}}, {"$set", bson.D{{"last_score_timestamp", userScore.LastScoreTimestamp}}}}).Decode(&user)
	if err != nil {
		c.JSON(404, gin.H{"msg": "user not found"})
		return
	}
	_, err = redisClient.ZIncrBy(ctx, "leaderboard", float64(userScore.Score), userScore.ID.Hex()).Result()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	_, err = redisClient.ZIncrBy(ctx, "leaderboard:"+strings.ToLower(user.Country), float64(userScore.Score), userScore.ID.Hex()).Result()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, gin.H{"msg": "Score submitted"})
}
