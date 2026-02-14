package util

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateRandomString() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return base32.StdEncoding.EncodeToString(bytes)
}
