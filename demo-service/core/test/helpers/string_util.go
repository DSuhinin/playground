package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

const (
	DefaultRandomStringLength = 16
)

//
// GenerateRandomString :: used in tests when need safe unique identifier;
// exact length is not really important, so using default.
//
func GenerateRandomString() string {
	return generateRandomStringOfSpecifiedLength(DefaultRandomStringLength)
}

//
// generateRandomStringOfSpecifiedLength :: creates random string of a specified length.
//
func generateRandomStringOfSpecifiedLength(length int) string {
	b := make([]byte, (length+1)/2)
	rand.Read(b)

	return hex.EncodeToString(b)[:length]
}
