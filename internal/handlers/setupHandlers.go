package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func SetupHandlers(app *fiber.App) {
	// Handle file uploads
	app.Post("/upload", uploadHandler)

	// Serve uploaded images
	app.Get("/public/uploads/:dir/:filename", getImageHandler)
}
