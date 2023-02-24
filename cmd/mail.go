package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type SendEmailRequest struct {
	Action      string `json:"Action"`
	Destination struct {
		ToAddresses  []string `json:"ToAddresses"`
		CcAddresses  []string `json:"CcAddresses"`
		BccAddresses []string `json:"BccAddresses"`
	} `json:"Destination"`
	Message struct {
		Body struct {
			Text struct {
				Data string `json:"Data"`
			} `json:"Text"`
			Html struct {
				Data string `json:"Data"`
			} `json:"Html"`
		} `json:"Body"`
		Subject struct {
			Data string `json:"Data"`
		} `json:"Subject"`
	} `json:"Message"`
	Source           string   `json:"Source"`
	ReplyToAddresses []string `json:"ReplyToAddresses"`
}

func sendEmail(c *gin.Context, dataDir string, logDir string) error {
	var request SendEmailRequest

	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		return err
	}

	// Mkdir dataDir and logDir
	err = os.Mkdir(dataDir, 0644)
	if err != nil {
		return err
	}

	err = os.Mkdir(logDir, 0644)
	if err != nil {
		return err
	}

	// Write html data to dataDir/body.html
	err = os.WriteFile(filepath.Join(logDir, "body.html"), []byte(request.Message.Body.Html.Data), 0644)
	if err != nil {
		return err
	}

	// Write body to dataDir/body.txt
	err = os.WriteFile(filepath.Join(logDir, "body.txt"), []byte(request.Message.Body.Text.Data), 0644)
	if err != nil {
		return err
	}

	// Write headers (Subject Data, ToAddress, CCAddresses, BccAddresses, ReplyToAddresses, Source) to dataDir/headers.txt
	headers := fmt.Sprintf("Subject: %s\nTo: %s\nCc: %s\nBcc: %s\nReply-To: %s\nFrom: %s\n",
		request.Message.Subject.Data,
		strings.Join(request.Destination.ToAddresses, ","),
		strings.Join(request.Destination.CcAddresses, ","),
		strings.Join(request.Destination.BccAddresses, ","),
		strings.Join(request.ReplyToAddresses, ","),
		request.Source,
	)
	err = os.WriteFile(filepath.Join(logDir, "headers.txt"), []byte(headers), 0644)
	if err != nil {
		return err
	}

	// Read file from templates/success.txt
	successTemplate, err := os.ReadFile("templates/success.txt")
	if err != nil {
		return err
	}

	// Replace {{message}} with absolute path of the body.html
	successMessage := strings.Replace(string(successTemplate), "{{message}}", filepath.Join(dataDir, "body.html"), -1)

	// Respond with the content & 200
	c.String(http.StatusOK, successMessage)

	return nil
}

func sendRawEmail(c *gin.Context, dateDir string, logFilePath string) {

}
