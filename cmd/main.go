package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RequestBody struct {
	Action string `json:"Action"`
}

func handler(c *gin.Context) {
	var reqBody RequestBody

	// Bind json
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Build dateDir
	dateTime := time.Now().Format("2006-01-02T15:04:05.000Z")
	dateDir := Config.OutputDir + "/" + dateTime[:10]
	fullDir := dateDir + "/" + dateTime[11:22] + ".log"

	// Actions
	switch reqBody.Action {
	case "SendEmail":
		sendEmail(c, dateDir, fullDir)
	case "SendRawEmail":
		sendRawEmail(c, dateDir, fullDir)
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unsupported action"})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func main() {
	// Read environment variables
	readConfigFromEnv()

	// Endpoints
	r := gin.Default()
	r.POST("/", handler)

	// Run
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
