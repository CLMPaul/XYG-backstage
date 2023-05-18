package tools

import (
	"math/rand"
	"strconv"
	"time"
)

func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code)
	return res
}
