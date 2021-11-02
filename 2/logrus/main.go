package main

import (
	"context"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	DatabaseURL = "postgres://sanyarise:123@127.0.0.1:5432/example?sslmode=disable"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("../mylog/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	a := app{}

	if err := a.Init(context.Background(), logger); err != nil {
		log.Fatal(err)
	}

	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}
}
