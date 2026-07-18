package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome the Golang App Development by sridhar")
	})
	fmt.Println("Server is running on port 5001")
	app.Listen(":5001") // had to change the port to 5001 since airplay is using 5000
	// now will start working from this project
	// Hope you're clear with intial setup

}
