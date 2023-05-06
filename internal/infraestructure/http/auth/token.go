package auth

import (
	config "backend-challenge-api/internal/infraestructure/configuration"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mockedKeySecret = []byte("SecretKey")

// Create an token that defines user permissions
func GenerateToken() (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["expt"] = time.Now().Add(time.Hour * 1).Unix()

	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return userToken.SignedString([]byte(config.GetSecretKey()))
}

// Verifies if token received in Request is valid
func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, tokenHasCorrespondentClaims := token.Claims.(jwt.MapClaims); tokenHasCorrespondentClaims && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, tokenHasTrueValue := token.Method.(*jwt.SigningMethodHMAC); !tokenHasTrueValue {
		return nil, fmt.Errorf("Unexpected signature method: %v\n", token.Header["alg"])
	}
	return config.GetSecretKey(), nil
}
