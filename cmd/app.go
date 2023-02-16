package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/mmatz101/go-odds/cmd/api"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Our website is working.")
	})
	app.Get("/:league/:week/:year/:season_type", vegasInsiderCall)
	app.Get("/retrieve/:league/:week/:year/:season_type", retrieveGames)
	app.Post("/users", api.CreateUser)
	app.Post("/login", api.LoginUser)
	app.Use("/", jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Post("/create-pick", api.CreatePick)
	app.Get("/auth/secure", api.SecureArea)
	app.Listen(":8000")
}
