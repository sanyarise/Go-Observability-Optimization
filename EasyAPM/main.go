package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))

	r.GET("/example", func(c *gin.Context) {
		span, ctx := apm.StartSpan(c.Request.Context(), "PingHandler", "request")
		defer span.End()

		processingRequest(ctx)

		todo, err := getTodoFromAPI(ctx)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})
	r.Run("0.0.0.0:8081")
}

func processingRequest(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "processingRequest", "custom")
	defer span.End()

	doSomething(ctx)

	// time sleep simulate some processing time
	time.Sleep(15 * time.Millisecond)
	return
}

func doSomething(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "doSomething", "custom")
	defer span.End()

	// time sleep simulate some processing time
	time.Sleep(20 * time.Millisecond)
	return
}

func getTodoFromAPI(ctx context.Context) (map[string]interface{}, error) {
	span, ctx := apm.StartSpan(ctx, "getTodoFromAPI", "custom")
	defer span.End()

	var result map[string]interface{}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, err
}
