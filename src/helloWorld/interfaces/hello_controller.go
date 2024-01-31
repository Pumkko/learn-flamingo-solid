package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type NicknameBody struct {
	Nickname string
}

type HelloController struct {
	responder *web.Responder
}

func (controller *HelloController) Inject(responder *web.Responder) *HelloController {
	controller.responder = responder
	return controller
}

func (controller *HelloController) ApiGreet(_ context.Context, r *web.Request) web.Result {
	nickname, _ := r.Params["nickname"]

	nicknameBody := NicknameBody{
		Nickname: nickname,
	}

	return controller.responder.Data(nicknameBody)
}

func (controller *HelloController) Index(_ context.Context, r *web.Request) web.Result {
	return controller.responder.Render("index", nil)
}
