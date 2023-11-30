package main

import (
	"crud_go/database"
	"crud_go/models"
	"crud_go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()
	database.DB.AutoMigrate(models.Users{})

	app := fiber.New()

	routes.UserRouter(app)

	log.Fatal(app.Listen(":3000"))
}
