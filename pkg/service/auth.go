package service

import (
	"crypto/sha1"
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
	userName string
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user jwt.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	authToken := token.NewWithClaims(token.SigningMethodHS256, &tokenClaims{
		token.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Username,
	})
	return authToken.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
