package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var Cost int

func HashPassword(password string) (string, error) {
	Cost = 12

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), 12)
	if hashErr != nil {
		return "", hashErr
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}
