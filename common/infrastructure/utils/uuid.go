package utils

import (
	uuid "github.com/satori/go.uuid"
	google_uuid "github.com/google/uuid"
)


// Generate a new UUID 
// using satori/go.uuid lib
func GetID() string {
	return uuid.NewV4().String()
}

// Simple validation of the provided UUID
func IsVaalidUUID(id string) error {
	_, err := google_uuid.Parse(id)
	return err
}
