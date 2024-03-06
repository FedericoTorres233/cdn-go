package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(file *multipart.FileHeader, filename string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

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
