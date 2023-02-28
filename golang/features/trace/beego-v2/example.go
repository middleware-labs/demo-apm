package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	mwbeego "github.com/middleware-labs/golang-apm-beego-beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	config, _ := track.Track(
		track.WithConfigTag("service", "beego-beego-pname"),
		track.WithConfigTag("projectName", "beego-beego-sname"),
	)
	web.Router("/", &MainController{})
	mware := mwbeego.MiddleWare(config.ServiceName)
	web.RunWithMiddleWares(":7777", mware)

}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString(fmt.Sprintf("hello world"))
}
