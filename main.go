package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/redisApi/handlers"
	"github.com/zeimedee/redisApi/services"
)

func handler(app *fiber.App) {

	app.Post("/login", handlers.Login)
	app.Get("/student", handlers.Student)
}

func main() {

	app := fiber.New()
	services.Connect()

	handler(app)

	log.Fatal(app.Listen(":4000"))

}
