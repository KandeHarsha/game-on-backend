package entities

import "github.com/golang-jwt/jwt/v5"

type AccessTokenJWT struct {
	jwt.RegisteredClaims
	SessionID string `json:"sid"`
}
