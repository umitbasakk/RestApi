package middleware

import "golang.org/x/crypto/bcrypt"

func ComparePassword(password string, hashPassword string) bool {

	res := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if res != nil {
		return false
	} else {
		return true
	}
}
