package auth

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email string, time time.Time) string {
	return email
}

func validPassword(password string) bool {
	length := len(password) >= 8
	caps := false
	lower := false
	number := false

	for _, char := range password {
		if strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(char)) {
			caps = true
		} else if strings.Contains("abcdefghijklmnopqrstuvwxyz", string(char)) {
			lower = true
		} else if strings.Contains("0123456789", string(char)) {
			number = true
		}
	}

	return length && caps && lower && number
}

func validEmail(email string) bool {
	at := false
	dot := false
	atPos := -1

	for i, char := range email {
		if char == '@' && i > 0 {
			if at {
				return false
			}
			at = true
			atPos = i
		} else if char == '.' && atPos+1 < i && i < len(email)-1 {
			if dot {
				return false
			}
			dot = true
		}
	}
	return at && dot
}
