package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mw_sql "github.com/middleware-labs/golang-apm-sql/sql"
	"github.com/middleware-labs/golang-apm/tracker"
	"net/http"
)

type User struct {
	ID    uint   `gorm:"primary_key"`
	Name  string
	Email string
}

// gormOpen is like gorm.Open, but it uses otelsql to instrument the database.
func gormOpen(driverName, dataSourceName string,dbName string) (*gorm.DB, error) {
	db, err := mw_sql.Open(driverName, dataSourceName,mw_sql.WithDBName(dbName),mw_sql.WithAttributes(tracker.String("db.system", "mysql")))
	if err != nil {
		return nil, err
	}
	conn,err := gorm.Open(driverName, db)
	return conn,err
}

func main()  {
	r := gin.Default()
	config, _ := tracker.Track(
		tracker.WithConfigTag("service", "Your service name"),
		tracker.WithConfigTag("projectName", "Your project name"),
		tracker.WithConfigTag("accessToken", "Your access token"),
	)
	r.Use(g.Middleware(config))
	db, err := gormOpen("mysql", "root:@tcp(127.0.0.1:3306)/next","next")
	if err != nil {
		fmt.Println("err",err)
	}
	r.GET("/todo", func(c *gin.Context) {
		user := User{Name: "John Doe", Email: "john@example.com"}
	    if err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email).Error; err != nil {
			fmt.Println("Error creating user:", err)
		     return
		}
		c.JSON(http.StatusOK, "ok")
	})
	_ = r.Run(":8081")
}

