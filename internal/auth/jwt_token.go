package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *AuthConfig) GenerateJWTToken(expiresAfter time.Duration, userID string) (string, error) {
	now := time.Now()
	after := now.Add(expiresAfter)

	JWTClaims := jwt.RegisteredClaims{
		Issuer:    "butchersorders",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(after),
		Subject:   userID,
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims)

	if newToken == nil {
		log.Println("Could not create a new Token")
		return "", errors.New("could not create a new Token")
	}

	signedToken, err := newToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		log.Printf("Could not sign the Token: %v", err)
		return "", errors.New("could not sign the Token")
	}
	return signedToken, nil
}
