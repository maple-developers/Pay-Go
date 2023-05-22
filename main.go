package main

import (
	/* 	"Pay-Go/database"*/
	"Pay-Go/database"
	"Pay-Go/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	fmt.Println("WE ARE HERE")
	const SecretKey = "w!z%C*F-JaNcRfUjXn2r5u8x/A?D(G+KbPeSgVkYp3s6v9y$B&E)H@McQfTjWmZq"
	/* fmt.Println("ApiKey", controllers.ApiKey("yaser@mapletele.com")) */
	const y int = 12
	database.ConnectDB()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
		}))
	/* app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "GO"})
	}) */
	routes.Setup(app)
	app.Listen(":80")

}
