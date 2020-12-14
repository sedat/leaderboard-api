package handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sedat/leaderboard-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLeaderboard(c *gin.Context) {

	val, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for _, v := range val {
		str, _ := v.Member.(string)
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		var user models.User
		findOptions := options.FindOne()
		findOptions.SetProjection(bson.M{"password": 0})
		err = usersCollection.FindOne(ctx, bson.M{"_id": oid}, findOptions).Decode(&user)
		rank, err := redisClient.ZRevRank(ctx, "leaderboard", user.ID.Hex()).Result()
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		user.Rank = int(rank + 1)
		users = append(users, user)
	}
	c.JSON(200, users)
}

func GetLeaderboardWithLimit(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(400, gin.H{"msg": "Limit must be a number"})
		return
	}
	val, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, int64(limit-1)).Result()
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for _, v := range val {
		str, _ := v.Member.(string)
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		var user models.User
		findOptions := options.FindOne()
		findOptions.SetSort(bson.M{"points": -1})
		findOptions.SetProjection(bson.M{"password": 0})
		err = usersCollection.FindOne(ctx, bson.M{"_id": oid}, findOptions).Decode(&user)
		rank, err := redisClient.ZRevRank(ctx, "leaderboard", user.ID.Hex()).Result()
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		user.Rank = int(rank + 1)
		users = append(users, user)
	}
	c.JSON(200, users)
}

func GetLeaderboardByCountry(c *gin.Context) {
	country := c.Param("country")

	val, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard:"+strings.ToLower(country), 0, -1).Result()
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for _, v := range val {
		str, _ := v.Member.(string)
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		var user models.User
		findOptions := options.FindOne()
		findOptions.SetSort(bson.M{"points": -1})
		findOptions.SetProjection(bson.M{"password": 0})
		err = usersCollection.FindOne(ctx, bson.M{"_id": oid}, findOptions).Decode(&user)
		rank, err := redisClient.ZRevRank(ctx, "leaderboard:"+strings.ToLower(country), user.ID.Hex()).Result()
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		user.Rank = int(rank + 1)
		users = append(users, user)
	}
	c.JSON(200, users)
}

func GetLeaderboardByCountryWithLimit(c *gin.Context) {
	country := c.Param("country")
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(400, gin.H{"msg": "Limit must be a number"})
		return
	}
	val, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard:"+strings.ToLower(country), 0, int64(limit-1)).Result()
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for _, v := range val {
		str, _ := v.Member.(string)
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		var user models.User
		findOptions := options.FindOne()
		findOptions.SetSort(bson.M{"points": -1})
		findOptions.SetProjection(bson.M{"password": 0})
		err = usersCollection.FindOne(ctx, bson.M{"_id": oid}, findOptions).Decode(&user)
		rank, err := redisClient.ZRevRank(ctx, "leaderboard:"+strings.ToLower(country), user.ID.Hex()).Result()
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		user.Rank = int(rank + 1)
		users = append(users, user)
	}
	c.JSON(200, users)
}

func GetLeaderboardInRange(c *gin.Context) {
	start, err := strconv.Atoi(c.Param("start"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	end, err := strconv.Atoi(c.Param("end"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	val, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard", int64(start), int64(end-1)).Result()
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	for _, v := range val {
		str, _ := v.Member.(string)
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		var user models.User
		err = usersCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
		rank, err := redisClient.ZRevRank(ctx, "leaderboard", user.ID.Hex()).Result()
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		user.Rank = int(rank + 1)
		users = append(users, user)
	}
	c.JSON(200, users)
}
