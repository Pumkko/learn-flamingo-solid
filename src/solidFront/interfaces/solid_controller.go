package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type SolidController struct {
	responder *web.Responder
}

func (controller *SolidController) Inject(responder *web.Responder) *SolidController {
	controller.responder = responder
	return controller
}

func (controller *SolidController) Index(_ context.Context, r *web.Request) web.Result {
	return controller.responder.Render("index", nil)
}
