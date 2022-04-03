package service

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/LuxAeterna-git/jwt"
	"github.com/LuxAeterna-git/jwt/pkg/repository"
	token "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "usedForDecodeToken"
	salt       = "justSalt"
)

type tokenClaims struct {
	token.StandardClaims
	userId string
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user jwt.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	accessToken := token.NewWithClaims(token.SigningMethodHS256, &tokenClaims{
		token.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID.String(),
	})
	return accessToken.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	access, err := token.ParseWithClaims(accessToken, &tokenClaims{}, func(t *token.Token) (interface{}, error) {
		if _, ok := t.Method.(*token.SigningMethodHMAC); !ok {
			return "", errors.New("Wrong token")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := access.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("Wrong token")
	}

	return claims.userId, nil

}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
