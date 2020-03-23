package model

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	jwt.StandardClaims
}

type contextKey string

const JwtClaimsCtxKey contextKey = "jwtClaims"
const JwtTokenKey contextKey = "jwtToken"
