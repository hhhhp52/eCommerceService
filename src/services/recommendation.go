package services

import (
	"context"
	"eCommerceService/src/db"
	"eCommerceService/src/db/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

func GetRecommendation(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
			return
		}
	}()
	// Check access token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token required"})
		return
	}

	tx := db.DB()
	// get user id from access token
	user, err := dao.GetUserByToken(tx, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
		return
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "docker.for.mac.localhost:6379", // Update with your Redis server address
	})
	ctx := context.Background()

	recommendationKey := strconv.Itoa(user.ID) + "_recommendations"

	// Check if the recommendations are in the cache
	cachedRecommendations, err := redisClient.Get(ctx, recommendationKey).Result()
	if err == nil {
		// Convert the recommendations to a map
		var recommendations []map[string]interface{}
		err = json.Unmarshal([]byte(cachedRecommendations), &recommendations)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unmarshal recommendations"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
		return
	}

	// If not, get the recommendations from the database (mocked here)
	recommendations, err := dao.GetRecommendationByUser(tx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get recommendations"})
		return
	}
	if recommendations == nil || len(recommendations) == 0 {
		c.JSON(http.StatusOK, gin.H{"recommendations": nil})
		return
	}

	// Convert the recommendations to a map
	recommendationsMap := make([]map[string]interface{}, 0)
	for _, recommendation := range recommendations {
		recommendationMap := map[string]interface{}{
			"id":    recommendation.ID,
			"name":  recommendation.Name,
			"price": recommendation.Price,
		}
		recommendationsMap = append(recommendationsMap, recommendationMap)
	}

	// Convert recommendationsMap to JSON
	recommendationsJSON, err := json.Marshal(recommendationsMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not marshal recommendations"})
		return
	}

	// Cache the recommendations
	err = redisClient.Set(ctx, recommendationKey, recommendationsJSON, 10*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cache recommendations"})
		return
	}

	// Return the recommendations
	c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
}
