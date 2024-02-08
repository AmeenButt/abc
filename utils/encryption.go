package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("unable to hash the password: %v", err)
	}

	return string(hashedPassword)
}

func CheckPasswd(pwdInDb string, userPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdInDb), []byte(userPwd))
	if err != nil {
		log.Fatalf("passwords do not match: %v", err)
		return false
	} else {
		return true
	}
}
