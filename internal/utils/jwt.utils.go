package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("secret-key")
var JWT_REFRESH_SECRET = []byte("refresh-secret")

/* FOR ACCESS TOKEN */
func GenerateJWT(uid int64) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{Subject: strconv.Itoa(int(uid)), ExpiresAt: jwt.NewNumericDate(exp)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_SECRET)
	return tokenString, err
}

/* FOR REFRESH TOKEN */
func GenerateRefreshToken(uid int64) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{Subject: strconv.Itoa(int(uid)), ExpiresAt: jwt.NewNumericDate(exp)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_SECRET)
	return tokenString, err
}
