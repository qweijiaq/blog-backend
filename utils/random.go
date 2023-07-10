package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v", rand.Intn(10000))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
