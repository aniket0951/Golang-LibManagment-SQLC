package utils

import (
	"fmt"
	"math/rand"
)

var latters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = latters[rand.Intn(len(latters))]
	}

	return string(b)
}

func RandomAuthor() string {

	return RandomString(6)
}

func RandomAddress() string {
	return RandomString(15)
}

func RandomBook() string {
	return RandomString(13)
}

func RandomEmail() string {
	rand_str := RandomString(13)
	return fmt.Sprintf("%s@gmail.com", rand_str)
}
