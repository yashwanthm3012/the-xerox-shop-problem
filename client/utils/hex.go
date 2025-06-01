package utils

import (
	"crypto/rand"
	"fmt"
)

// Generate4DigitHex generates a random 4-digit hexadecimal string (e.g., "1A3F").
func Generate4DigitHex() (string, error) {
	var b [2]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	num := uint16(b[0])<<8 | uint16(b[1])
	return fmt.Sprintf("%04X", num), nil
}
