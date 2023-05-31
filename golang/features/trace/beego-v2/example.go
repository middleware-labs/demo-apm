package main

import (
	"fmt"
	// "github.com/pyroscope-io/client/pyroscope"
	"github.com/beego/beego/v2/server/web"
	mwbeego "github.com/middleware-labs/golang-apm-beego-beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag("service", "beego-v2-app-legacy"),
		track.WithConfigTag("projectName", "beego-v2-keval-project"),
		track.WithConfigTag("accessToken", "Your access token"),
	)
	web.Router("/v24", &MainController{})
	mware := mwbeego.MiddleWare(config.ServiceName)
	web.RunWithMiddleWares(":7777", mware)

}
// MW_PROFILING_SERVER_URL=https://profiling.front.env.middleware.io go run example.go

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString(fmt.Sprintf("hello world"))
}
