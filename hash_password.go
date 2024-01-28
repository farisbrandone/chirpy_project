package main

import (
	"golang.org/x/crypto/bcrypt"
)

func GenereCriptPassword (param string) (string, error) {
	bcryptPassword,err:=bcrypt.GenerateFromPassword([]byte(param),14)
	return string(bcryptPassword),err
}

