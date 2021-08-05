package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yafiesetyo/merempah-api-clone/server/handlers"
)

// ProductRoute...
type ProductRoute struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute
func (route ProductRoute) RegisterRoute() {
	handler := handlers.ProductHandler{Handler: route.Handler}
	r := route.RouterGroup.Group("/api/product")

	r.Get("", handler.SelectAll)
	r.Get("/id/:id", handler.FindById)
	r.Post("", handler.Add)
	r.Put("/id/:id", handler.Edit)
	r.Delete("/id/:id", handler.Delete)
}
