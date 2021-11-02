package main

import (
	"context"
	"log"
	"go.uber.org/zap"
)

const (
	DatabaseURL = "postgres://sanyarise:123@localhost:5432/example?sslmode=disable"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = logger.Sync()
	}()

	a := app{}

	if err := a.Init(context.Background(), logger); err != nil {
		log.Fatal(err)
	}

	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}
}