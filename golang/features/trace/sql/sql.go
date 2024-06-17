package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	mw_sql "github.com/middleware-labs/golang-apm-sql/sql"
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
	r.GET("/test", func(c *gin.Context) {
		ctx := c.Request.Context()
		db, err := mw_sql.Open("mysql", "uname:password@tcp(127.0.0.1:3306)/todo")
		if err != nil {
			track.RecordError(ctx, err)
		}
		var name string
		err = db.QueryRowContext(ctx, "select p.title from tutorials as p where p.id = :id;", sql.Named("id", 1)).Scan(&name)
		if err != nil {
			track.RecordError(ctx, err)
		}
		c.JSON(http.StatusOK, "ok")
	})
	_ = r.Run(":8081")
}
