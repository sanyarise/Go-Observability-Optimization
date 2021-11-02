package main

import (
	"context"
	"log"
	"github.com/sirupsen/logrus"
)

const (
	DatabaseURL = "postgres://sanyarise:123@127.0.0.1:5432/example?sslmode=disable"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	a := app{}

	if err := a.Init(context.Background(), logger); err != nil {
		log.Fatal(err)
	}

	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}
}
