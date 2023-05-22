package routes

import (
	"Pay-Go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//const SecretKey = "w!z%C*F-JaNcRfUjXn2r5u8x/A?D(G+KbPeSgVkYp3s6v9y$B&E)H@McQfTjWmZq"
	app.Get("/", controllers.HomeStatus)
	app.Get("/Account", controllers.Account)
	app.Get("/CreateOrder", controllers.CreateOrder)
	app.Get("/ListTickerPrices", controllers.ListTickerPrices)
	app.Get("/ListOrder", controllers.ListOrders)
	app.Get("/C2CTradeHistory", controllers.C2CTradeHistory)
}
