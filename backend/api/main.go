package main

import (
	"Server/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/sridharakella/golangapp/backend/docs"
)

// @title Golang App Development by sridhar
// @version 1.0
// @description This is a sample server celler by sridhar
// @host localhost:5001
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type into the textbox: Bearer {your JWT token}

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		//AllowOrigins: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome the Golang App Development by sridhar")
	})
	fmt.Println("Server is running on port 5001")

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":5001") // had to change the port to 5001 since airplay is using 5000
	// now will start working from this project
	// Hope you're clear with intial setup

}
