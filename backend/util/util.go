package util

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func GenerateSessionToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Error generating session token: %v", err)
	}
	return hex.EncodeToString(b)
}