package main

import (
	"fmt"

	"github.com/federicotorres233/cdn-fiber/handlers"
	"github.com/gofiber/fiber/v3"
)

const port = 3000

func main() {
	app := fiber.New()

	// Handle file uploads
	app.Post("/upload", handlers.UploadHandler)

	// Serve uploaded images
	app.Get("/image/:dir/:filename", handlers.GetImageHandler)

	// Start the server
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}

}
