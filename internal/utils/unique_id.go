package utils

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

//goland:noinspection GoUnusedExportedFunction
func UUID() string {
	return uuid.New().String()
}

//goland:noinspection GoUnusedExportedFunction
func UUIDWithoutDash() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

//goland:noinspection GoUnusedExportedFunction
func XID() string {
	return xid.New().String()
}

func RandomBytes(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	rand.Read(bytes)
	return bytes
}

// 字母和数字，移除了 oODlL01 等易混淆字母
var randomStringLetters = []rune("abcdefghijkmnpqrstuvwxyzABCEFGHIJKMNPQRSTUVWXYZ23456789")

// RandomString returns a random string with a fixed length
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randomStringLetters[rand.Intn(len(randomStringLetters))]
	}

	return string(b)
}

// Nonce returns nonce string, param `size` better for even number.
func Nonce(size uint) string {
	nonce := make([]byte, size/2)
	io.ReadFull(cryptoRand.Reader, nonce)

	return hex.EncodeToString(nonce)
}
