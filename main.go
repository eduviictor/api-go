package main

import (
	"api-go/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func initializeRoutes(app *fiber.App) {
	app.Get("/", routes.HelloWorld)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Post("/users", routes.CreateUser)
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)
}

func main() {
		app := fiber.New()

		initializeRoutes(app)

		err := godotenv.Load()

		if err != nil {
			log.Fatal("Erro loading .env file")
		}

		port := os.Getenv("PORT")

		app.Listen(":" + port)
}