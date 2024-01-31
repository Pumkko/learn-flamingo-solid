package solidfront

import (
	"pumkko/learnFlamingo/src/solidFront/interfaces"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct {
}

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	solidController *interfaces.SolidController
}

func (r *routes) Inject(controller *interfaces.SolidController) *routes {
	r.solidController = controller
	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/", "index")
	registry.HandleGet("index", r.solidController.Index)

	registry.MustRoute("/*name", `flamingo.static.file(name, dir?="frontend/dist")`)
}
