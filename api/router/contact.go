package router

import (
	"github.com/gofiber/fiber/v2"

	"srv_order/main/api/handlers"
	contact "srv_order/main/pkg/order"
)

func ContactRouter(app fiber.Router, service contact.Service) {
	app.Get("/orders", handlers.Gets(service))
	app.Post("/orders", handlers.AddOrder(service))
}
