package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	UserID    string
	UserRole  string
	UserEmail string
	jwt.StandardClaims
}

type ContextUserData struct {
	UserID    string
	UserRole  string
	UserEmail string
}

type contextKey string

const ContextKey = contextKey("UserData")
