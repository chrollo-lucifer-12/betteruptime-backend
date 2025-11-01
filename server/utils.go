package server

import (
	"crypto/rand"
	"crypto/sha256"
)

func GenerateSecureRandomString() (string, error) {
	const alphabet = "abcdefghijkmnpqrstuvwxyz23456789"
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]byte, len(bytes))
	for i, b := range bytes {
		id[i] = alphabet[b>>3]
	}

	return string(id), nil
}

func HashSecret(secret string) []byte {
	buff := sha256.Sum256([]byte(secret))
	return buff[:]
}
