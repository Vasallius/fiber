package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = ":3000"
	// } else {
	// 	port = ":" + port
	// }
	// fmt.Println(port)
	// return port
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HELLO WORLD PLEASE WORK",
		})
	})
	fmt.Println("test")
	SayHello()
	app.Listen("3000")
}
