package main

import (
	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	mw_gorm "github.com/middleware-labs/go-agent-gorm/gorm"
	"github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

type User struct {
	ID    uint   `gorm:"primary_key"`
	Name  string
	Email string
}

func main() {
	r := gin.Default()
	config, _ := tracker.Track(
		tracker.WithConfigTag("service", "Your service name"),
		tracker.WithConfigTag("projectName", "Your project name"),
		tracker.WithConfigTag("accessToken", "your access token"),
	)
	r.Use(g.Middleware(config))
	r.GET("/todo", func(c *gin.Context) {
		ctx := c.Request.Context()
		db, err := gorm.Open(mysql.Open("username:password@tcp(127.0.0.1:3306)/dbname"), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "db connection error")
			return
		}
		p := mw_gorm.NewPlugin(mw_gorm.WithDBName("dbname"), mw_gorm.WithAttributes(tracker.String("db.system", "mysql")))
		if err := db.Use(p); err != nil {
			c.JSON(http.StatusInternalServerError, "db connection error")
			return
		}
		user := User{Name: "John Doe", Email: "john@example.com"}
		if err := db.WithContext(ctx).Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
		c.JSON(http.StatusOK, "ok")
	})
	_ = r.Run(":8081")
}


