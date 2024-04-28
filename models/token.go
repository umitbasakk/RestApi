package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Userid   int
	Username string
	jwt.StandardClaims
}
