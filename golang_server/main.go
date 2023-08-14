package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Player struct {
	Username string `json:"username"`
	Damage   string `json:"damage"`
}

// Simulated user database for our POC
var users = map[int]string{
	1: "Alice",
	2: "Bob",
	3: "Charlie",
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const validToken = "123-321-567-765"

		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		// Splitting the tokenHeader based on the " " (space) delimiter
		splitToken := strings.Split(tokenHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token format"})
			c.Abort()
			return
		}

		tokenValue := splitToken[1]

		// Assuming validToken is the string containing the correct token
		if tokenValue != validToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserHandler(c *gin.Context) {
	// Extracting userID from the query parameters
	userIDStr := c.DefaultQuery("userID", "0")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID format"})
		return
	}

	// Fetch user from the simulated database
	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": userID, "username": user})
}

func GroupByDamage(c *gin.Context) {
	var data struct {
		Players []Player `json:"players"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Bad request data"})
		return
	}

	grouped := make(map[string][]Player)

	for _, player := range data.Players {
		key := "lobby" + player.Damage + "Players"
		grouped[key] = append(grouped[key], player)
	}

	c.JSON(200, grouped)
}

func main() {
	r := gin.Default()
	// This r.use token.. applies the requirement of a token to all the routes
	r.Use(TokenAuthMiddleware())
	r.GET("/getuser", GetUserHandler)
	r.POST("/group", GroupByDamage)
	r.Run(":8080")
}
