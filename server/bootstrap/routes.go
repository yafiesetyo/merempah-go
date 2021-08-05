package bootstrap

import (
	"github.com/yafiesetyo/merempah-api-clone/server/bootstrap/routes"
	"github.com/yafiesetyo/merempah-api-clone/server/handlers"
)

func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
	}

	apiV1 := boot.App.Group("/v1")

	productRoutes := routes.ProductRoute{RouterGroup: apiV1, Handler: handler}
	productRoutes.RegisterRoute()
}
