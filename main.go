package main

import "golang.org/x/crypto/bcrypt"

func main() {
	data, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	println(string(data))
}
