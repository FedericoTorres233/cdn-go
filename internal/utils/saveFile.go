package utils

import (
	"io"
	"mime/multipart"
	"os"
)

// SaveFile saves a file to disk and name it to a certain filename
func SaveFile(file *multipart.FileHeader, filename string) error {

  // Open file
  src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
  
  // Create file on disk
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
