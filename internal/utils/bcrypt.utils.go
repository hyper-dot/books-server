package utils

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

/* GENERATES RANDOM SALT FOR EACH USERS */
func generateRandomSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password string) (string, string, error) {
	salt, err := generateRandomSalt(10)
	passwordBytes := []byte(salt + password)
	hashedPass, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return "", "", err
	}
	return string(hashedPass), salt, nil
}

func CheckPassword(password string, hash string) bool {
	hashedPassword := []byte(hash)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	return err == nil
}
