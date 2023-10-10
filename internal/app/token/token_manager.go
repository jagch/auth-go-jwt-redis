package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	Create(userID string) (string, error)
	Verify(token string) (bool, error)
	Renew(token string) (string, error)
}

type TokenJWT struct {
	secretKey []byte
}

func NewTokenJWT(secretKey string) TokenJWT {
	return TokenJWT{
		secretKey: []byte(secretKey),
	}
}

func (t TokenJWT) Create(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t TokenJWT) Verify(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method: %v", token.Header["alg"])
		}

		return t.secretKey, nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (t TokenJWT) Renew(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method: %v", token.Header["alg"])
		}

		return t.secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token, it can't renew")
	}

	// Verify the expiration and renew the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not obtain token claims")
	}

	// New expiration
	expirationTime := time.Now().Add(time.Hour * 2)
	claims["exp"] = expirationTime.Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = newToken.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
