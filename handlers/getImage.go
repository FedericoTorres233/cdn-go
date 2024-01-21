package handlers

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func GetImageHandler(c fiber.Ctx) error {
	// Get the filename parameter from the URL
	filename := c.Params("filename")
	dir := c.Params("dir")

	// Construct the full path to the image
	filepath := uploadDir + dir + "/" + filename

	// Check if the file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
    return c.SendFile("./images/dummy.png")
    //return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Image not found"})
	}

	// Serve the image using Fiber's SendFile function
	return c.SendFile(filepath)
}
