package util

import (
	"crypto/rand"
	"encoding/base32"
)


func generateRandomString() string {
	bytes := make([]byte, 12)
	rand.Read(bytes)
	return base32.StdEncoding.EncodeToString(bytes)
}