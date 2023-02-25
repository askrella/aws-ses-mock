package main

import (
	"github.com/askrella/ses-mock/internal"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("AWS SES Mock by Askrella Software Agency")
	internal.StartServer()
}
