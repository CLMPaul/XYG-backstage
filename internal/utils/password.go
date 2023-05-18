package utils

import (
	"crypto/sha256"
	"crypto/subtle"

	"golang.org/x/crypto/pbkdf2"
)

const PasswordSaltLength = 16
const PasswordHashLength = 32
const PasswordIterations = 10_0000

//goland:noinspection GoUnusedExportedFunction
func GeneratePassowrdHash(password string) ([]byte, []byte) {
	salt := RandomBytes(PasswordSaltLength)
	hash := pbkdf2.Key([]byte(password), salt, PasswordIterations, PasswordHashLength, sha256.New)
	return hash, salt
}

//goland:noinspection GoUnusedExportedFunction
func CheckPasswordHash(password string, hash, salt []byte) bool {
	expected := pbkdf2.Key([]byte(password), salt, PasswordIterations, PasswordHashLength, sha256.New)
	return subtle.ConstantTimeCompare(expected, hash) > 0
}
