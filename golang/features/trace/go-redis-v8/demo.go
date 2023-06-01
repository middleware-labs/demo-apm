package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	mwredis "github.com/middleware-labs/go-redis-v8/redis"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	track "github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

func FindBooks(c *gin.Context) {
	ctx := c.Request.Context()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (leave empty if not set)
		DB:       0,                // Redis database number
	})
	rdb.AddHook(mwredis.NewTracingHook())
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
	)
	r.Use(g.Middleware(config))
	r.GET("/books", FindBooks)
	r.Run(":8092")
}