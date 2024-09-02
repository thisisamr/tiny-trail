package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thisisamr/tiny-trail/Middleware"
	"github.com/thisisamr/tiny-trail/db"
	"github.com/thisisamr/tiny-trail/routes"
)

func make_routes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ClipUrl)
}

// main
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.CreatredisClient(0)
	app := fiber.New()

	app.Use(MiddleWare.Rate_Limiter)
	app.Use(logger.New())
	// register the routes
	make_routes(app)

	listen_error := app.Listen(":3000")
	if listen_error != nil {
		log.Fatal(err)
	}
	defer db.Redis_Client.Close()
}
