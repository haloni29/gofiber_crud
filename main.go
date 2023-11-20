package main

import (
	"crud_go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.UserRouter(app)

	log.Fatal(app.Listen(":3000"))
}
