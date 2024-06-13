package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	mw_pg "github.com/middleware-labs/golang-apm-pg/pg"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	r := gin.Default()
	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
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
