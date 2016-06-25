package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func Compare(raw, encryptedPassword string) bool {
	log.Println(raw + ":" + encryptedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(raw))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
