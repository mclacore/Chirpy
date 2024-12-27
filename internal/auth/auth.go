package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost, costErr := bcrypt.Cost([]byte(password))
	if costErr != nil {
		return "", costErr
	}

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), cost)
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
