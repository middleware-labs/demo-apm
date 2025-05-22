package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	// Sentry SDK
	// "github.com/getsentry/sentry-go"
	// sentrygin "github.com/getsentry/sentry-go/gin"

	// Middleware SDK
	g "github.com/middleware-labs/golang-apm-gin/gin"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	r := gin.Default()

	// if err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://bad64c222761bc2017f1e3e852bf47ed@o4508103606140928.ingest.us.sentry.io/4509292469878784",
	// }); err != nil {
	// 	fmt.Printf("Sentry initialization failed: %v\n", err)
	// }

	config, _ := track.Track(
		track.WithConfigTag("service", "golang-opsai"),
		// track.WithConfigTag("accessToken", "lneevcqdcnzccpdpbhelkxpjffdxeznkryvb"),
		// track.WithConfigTag("target", "ruplp.middleware.io:443"),

		// github-ai stage account
		track.WithConfigTag("accessToken", "mgqjtlgshkyhlykoaermkzbjpmgprkrzmbsb"),
		track.WithConfigTag("target", "sbncr.stage.env.middleware.io:443"),
		track.WithConfigTag(track.Debug, true),
		track.WithConfigTag(track.PauseMetrics, true),
	)

	r.Use(g.Middleware(config))
	r.Use(CustomRecovery())
	// r.Use(sentrygin.New(sentrygin.Options{}))

	r.GET("/booklist", FindBooks)
	r.GET("/booklistexception", FindBooksException)
	r.GET("/book-catalog", BookCatalog)
	r.GET("/show-user", ShowUser)
	r.Run(":8090")
}

func FindBooks(c *gin.Context) {
	// Simulate a genuine error, such as a database query failure
	test := []string{"test"}

	fmt.Println(test[3]) // This will cause a panic: index out of range

}

func FindBooksException(c *gin.Context) {
	// Simulate a panic to generate an error
	fmt.Println("Simulating a panic to generate an error")
	// Simulate a genuine error, such as a database query failure
	test := []string{"test"}
	fmt.Println(test[3]) // This will cause a panic: index out of range
	c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})

}

type User struct {
	Name string
}

func printUserName(u *User) {
	fmt.Println("User name is:", u.Name)
}

func ShowUser(c *gin.Context) {
	var user *User // user is nil
	printUserName(user)
}

func BookCatalog(c *gin.Context) {
	targetJson := `{"title": "Book 1", "author": "Author 1",}`
	var book map[string]interface{}
	json.Unmarshal([]byte(targetJson), &book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func FindBooksBr(c *gin.Context) {
	// This will cause a nil pointer dereference (genuine-looking bug)
	var bookList *[]string
	// Attempt to access the first element of a nil slice pointer
	firstBook := (*bookList)[0]
	fmt.Println("First book:", firstBook)
	c.JSON(http.StatusOK, gin.H{"book": firstBook})
}

// CustomRecovery returns a Gin middleware that records stack traces on panics
func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// 1. Fetch the current span
		span := trace.SpanFromContext(c.Request.Context())

		// 2. Preserve original error type
		var err error
		if recErr, ok := recovered.(error); ok {
			err = recErr
		} else {
			err = fmt.Errorf("%v", recovered)
		}

		// 3. Record error with stack trace
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())

		// 4. Send HTTP 500 response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
	})
}
