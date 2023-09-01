package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	mw_pg "github.com/middleware-labs/golang-apm-pg/pg"
	track "github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

func main() {
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag("service", "Your service name"),
		track.WithConfigTag("projectName", "Your project name"),
		track.WithConfigTag("accessToken", "your access token"),
	)
	r.Use(g.Middleware(config))

	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "dbname",
	})
	mw_pg.AddQueryHook(db)
	defer db.Close()
	r.GET("/books", func(c *gin.Context) {
		ctx := c.Request.Context()
		if err := db.Ping(ctx); err != nil {
			track.RecordError(ctx, err)
			panic(err)
		}
		c.JSON(http.StatusOK, "ok")
	})
	_ = r.Run(":8081")
}
