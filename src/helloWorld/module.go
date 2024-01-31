package helloworld

import (
	"pumkko/learnFlamingo/src/helloWorld/interfaces"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct {
}

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	helloController *interfaces.HelloController
}

func (r *routes) Inject(controller *interfaces.HelloController) *routes {
	r.helloController = controller
	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/api/greet/:nickname", "api.greet")
	registry.HandleGet("api.greet", r.helloController.ApiGreet)

	registry.MustRoute("/", "index")
	registry.HandleGet("index", r.helloController.Index)

	registry.MustRoute("/*name", `flamingo.static.file(name, dir?="frontend/dist")`)

}
