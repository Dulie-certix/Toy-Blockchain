package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA-256 Hashing function
func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}