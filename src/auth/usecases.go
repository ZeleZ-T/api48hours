package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString([]byte("secret key"))

	return tokenString, err
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(jwt.SigningMethod); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret key"), nil
	})
	if err != nil {
		return "", err
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); ok {
		return claims["email"].(string), nil
	}
	return "", errors.New("unexpected")
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
