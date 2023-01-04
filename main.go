package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var password string
	var cost int

	flag.StringVar(&password, "p", "verysecret", "a plain password")
	flag.IntVar(&cost, "c", 12, "a password's cost")

	flag.Parse()

	fmt.Println("Generating hash for password:", password, "with cost", cost, "\n")

	passwordByte := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, cost)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(hashedPassword))
}
