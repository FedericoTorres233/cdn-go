package handlers

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/federicotorres233/cdn-fiber/internal/types"
	"github.com/federicotorres233/cdn-fiber/internal/utils"
	"github.com/gofiber/fiber/v3"
)

const uploadDir = "./public/uploads/"

func uploadHandler(c fiber.Ctx) error {
	// Parse the form data, including files
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Create the uploads directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Iterate through uploaded files
	file := form.File["file"][0]
	metadata_json := form.Value["metadata"][0]
	var metadata types.Metadata
	err1 := json.Unmarshal([]byte(metadata_json), &metadata)
	if err1 != nil {
		log.Println("Error parsing JSON:", err1)
		return err1
	}
	log.Println(metadata)
	log.Println(form)

	// Generate UUID
	uuid := utils.CreateUUID() + filepath.Ext(file.Filename)

	// Generate a unique filename
	filename := filepath.Join(uploadDir+"/"+metadata.Dir, uuid)

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(uploadDir+"/"+metadata.Dir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Save the file to disk
	if err := utils.SaveFile(file, filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Resize original image
	if err := utils.ResizeImage(filename, "", metadata.Size); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Resize image more scales
	for tag, size := range metadata.Scales {
		if err := utils.ResizeImage(filename, tag, size); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "Files uploaded successfully", "uuid": uuid})
}
