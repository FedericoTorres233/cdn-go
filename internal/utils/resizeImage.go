package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ResizeImage(filename string, tag string, scale int) error {

	// Extract the file extension
	ext := filepath.Ext(filename)

	if scale == 0 {
		return nil
	}

	// Generate the output filename without extension
	var outFilename string
	if tag == "" {
		outFilename = filename[:len(filename)-len(ext)] + "_original" + ext
	} else {
		outFilename = filename[:len(filename)-len(ext)] + "_" + tag + ext
	}

	log.Println(filename)
	log.Println(outFilename + ".resized")
	log.Println(fmt.Sprintf("scale=%v:-1", scale))

	// Construct the FFmpeg command to resize the image
	cmd := exec.Command("ffmpeg", "-i", filename, "-vf", fmt.Sprintf("scale=%v:-1", scale), outFilename)

	// Run the FFmpeg command
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Rename file if original
	if tag == "" {
		err1 := os.Rename(outFilename, filename)
		if err1 != nil {
			return err1
		}
	}

	return nil
}
