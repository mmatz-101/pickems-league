package utils

import (
	"math/rand"
	"time"
)

// RandomStringGenerator creates of a string of random alphebetical character of n length
func RandomStringGenerator(n int) string {
	rand.Seed(time.Now().Unix())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
