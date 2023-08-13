package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/blazee5/finance-tracker/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	salt       = "sDFJKldsfkjFllj"
	signingKey = "shjahJASFJHadshio*sd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId primitive.ObjectID `json:"user_id"`
}

func (s *Service) CreateUser(user models.User) (interface{}, error) {
	user.Password = GenerateHashPassword(user.Password)
	id, err := s.Storage.UserDAO.Create(context.Background(), user)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (s *Service) GenerateToken(email, password string) (string, error) {
	user, err := s.Storage.UserDAO.GetUser(context.Background(), email, GenerateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(tokenTTL),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func GenerateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
