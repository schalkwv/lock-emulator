package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schalkwv/lock-emulator/database"
	"github.com/schalkwv/lock-emulator/routes"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
}

func main() {
	//handler := http.HandlerFunc(LockServer)
	//log.Fatal(http.ListenAndServe(":5000", handler))

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("welcome!")
}
