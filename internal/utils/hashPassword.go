package utils

import (
	"bosco-backend/internal/constants"
	"crypto/sha1"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(constants.SALT)))
}
