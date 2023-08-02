package config

import "github.com/golang-jwt/jwt/v4"

var JWTKey = []byte("secret")

type JwtClaims struct {
	Username string
	jwt.RegisteredClaims
}
