package main

import (
	"Server/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// @title Golang App Development by sridhar
// @version 1.0
// @description This is a sample server celler by sridhar
// @host localhost:5001
// @BasePath /

func main() {
	database.Connect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome the Golang App Development by sridhar")
	})
	fmt.Println("Server is running on port 5001")
	app.Listen(":5001") // had to change the port to 5001 since airplay is using 5000
	// now will start working from this project
	// Hope you're clear with intial setup

}
