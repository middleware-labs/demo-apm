package main

import (
	"github.com/beego/beego/v2/server/web"
	mwbeego "github.com/middleware-labs/golang-apm-beego-beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {

	ctrl := &MainController{}

	web.Router("/", ctrl)

	// GET http://localhost:8080/health => ctrl.Health()
	web.Router("/health", ctrl, "get:Health")

	// POST http://localhost:8080/update => ctrl.Update()
	web.Router("/update", ctrl, "post:Update")

	// support multiple http methods.
	// POST or GET http://localhost:8080/update => ctrl.GetOrPost()
	web.Router("/getOrPost", ctrl, "get,post:GetOrPost")

	// support any http method
	// POST, GET, PUT, DELETE... http://localhost:8080/update => ctrl.Any()
	web.Router("/any", ctrl, "*:Any")
	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
	)
	mware := mwbeego.MiddleWare(config.ServiceName)
	web.RunWithMiddleWares(":8080", mware)
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

// address: http://localhost:8080 GET
func (ctrl *MainController) Get() {

	// views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Get()"

	// don't forget this
	_ = mwbeego.Render(&ctrl.Controller)
}

// GET http://localhost:8080/health
func (ctrl *MainController) Health() {
	// views/hello_world.html
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "Health()"

	// don't forget this
	_ = mwbeego.Render(&ctrl.Controller)
}

// POST http://localhost:8080/update
func (ctrl *MainController) Update() {
	// views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "Update()"

	// don't forget this
	_ = mwbeego.Render(&ctrl.Controller)
}

// GET or POST http://localhost:8080/update
func (ctrl *MainController) GetOrPost() {
	// views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "GetOrPost()"

	// don't forget this
	_ = mwbeego.Render(&ctrl.Controller)
}

// any http method http://localhost:8080/any
func (ctrl *MainController) Any() {
	// views/hello_world.html
	ctrl.TplName = "hello_world.html"

	ctrl.Data["name"] = "Any()"

	// don't forget this
	_ = mwbeego.Render(&ctrl.Controller)
}
