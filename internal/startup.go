package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestBody struct {
	Action string `json:"Action"`
}

func handler(c *gin.Context) {
	var reqBody RequestBody

	// Bind json
	err := c.ShouldBindBodyWith(&reqBody, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Build dateDir
	dateTime := time.Now().Format("2006-01-02T15:04:05.000Z")
	dateDir := Config.OutputDir + "/" + dateTime[:10]
	logDir := dateDir + "/" + dateTime[11:22] + "-log"

	// Actions
	switch reqBody.Action {
	case "SendEmail":
		mailErr := SendEmail(c, dateDir, logDir)

		if mailErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": mailErr.Error(),
			})
			return
		}

		break
	case "SendRawEmail":
		SendRawEmail(c, dateDir, logDir)

		break
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "unsupported action"})
		return
	}
}

func StartServer() {
	logrus.Info("Starting mock server under port 8080")
	// Read environment variables
	ReadConfigFromEnv()

	// Endpoints
	r := gin.Default()
	r.POST("/", handler)

	// Run
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
