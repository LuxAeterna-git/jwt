package service

import (
	"github.com/LuxAeterna-git/jwt"
	"github.com/LuxAeterna-git/jwt/pkg/repository"
)

type Authorization interface {
	CreateUser(user jwt.User) (string, error)
	GenerateToken(username, password string) (string, error)
	GenerateRefreshToken() (string, error)
	ParseToken(token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repo.Authorization)}
}
