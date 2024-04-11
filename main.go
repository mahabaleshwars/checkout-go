package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	r := rand.NewSource(time.Now().UnixMilli())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var passwd string
	for i := 0; i < 8; i++ {
		passwd += string(chars[r.Int63()%int64(len(chars))])
	}

	data, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	println(string(data))
}
