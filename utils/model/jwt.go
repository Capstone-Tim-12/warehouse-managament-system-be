package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId   int    `json:"userId"`
	UserRole string `json:"userRole"`
	jwt.RegisteredClaims
}
