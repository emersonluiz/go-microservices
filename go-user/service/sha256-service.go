package service

import (
	"crypto/sha256"
	"fmt"
)

func Encoder(word string) string {
	encoded := sha256.Sum256([]byte(word))

	return fmt.Sprintf("%s", encoded)
}
