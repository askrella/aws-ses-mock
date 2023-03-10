package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RequestBody struct {
	Action string `json:"Action"`
}

func handler(c *gin.Context) {
	var reqBody RequestBody

	// Read the request body as a string
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	bodyString := string(bodyBytes)

	values, _ := url.ParseQuery(bodyString)
	reqBody = RequestBody{Action: values.Get("Action")}

	fmt.Println(reqBody) // prints the decoded request

	// Build dateDir
	dateTime := time.Now().Format("2006-01-02-15-04-05.000Z")
	dateDir := Config.OutputDir + "/" + dateTime[:10]
	logDir := dateDir + "/" + dateTime[11:22] + "-log"

	// Actions
	switch reqBody.Action {
	case "SendEmail":
		mailErr := SendEmail(bodyString, c, dateDir, logDir)

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
	// Read environment variables
	ReadConfigFromEnv()
	logrus.Info("Starting mock server under port ", Config.Port)

	// Endpoints
	r := gin.Default()
	r.POST("/", handler)

	// Run
	err := r.Run(":" + strconv.Itoa(Config.Port))
	if err != nil {
		panic(err)
	}
}
