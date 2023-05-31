package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mwredis "github.com/middleware-labs/go-redis-v9/redis"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	track "github.com/middleware-labs/golang-apm/tracker"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func FindBooks(c *gin.Context) {
	ctx := c.Request.Context()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (leave empty if not set)
		DB:       0,                // Redis database number
	})
	if err := mwredis.InstrumentTracing(rdb); err != nil {
		panic(err)
	}
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func main() {
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
		track.WithConfigTag("accessToken", "Your access token"),
	)
	r.Use(g.Middleware(config))
	r.GET("/books", FindBooks)
	r.Run(":8092")
}
