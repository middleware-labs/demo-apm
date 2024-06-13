package main

import (
	"github.com/astaxie/beego"
	mw_beego "github.com/middleware-labs/golang-apm-beego/beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)

type exampleController struct {
	beego.Controller
}

func (c *exampleController) Get() {
	c.Ctx.WriteString("Hello, world!")
}

func (c *exampleController) Template() {
	c.TplName = "hello.tpl"
}

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
	)
	// To enable tracing on template rendering, disable autorender
	beego.BConfig.WebConfig.AutoRender = false
	beego.Router("/hello", &exampleController{})
	beego.Router("/", &exampleController{}, "get:Template")
	mware := mw_beego.Middleware(config)
	beego.RunWithMiddleWares(":7777", mware)

}
