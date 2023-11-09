package auth

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt       = "sDFJKldsfkjFllj"
	signingKey = "shjahJASFJHadshio*sd"
	tokenTTL   = 12 * time.Hour
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userId,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return "", err
	}

	return claims.UserId, nil
}

func GenerateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
