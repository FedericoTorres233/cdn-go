package utils

import (
	"github.com/google/uuid"
)

func CreateUUID() string {
	// Generate a new UUID
	newUUID := uuid.New()

	return newUUID.String()
}
