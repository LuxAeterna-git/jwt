package service

import "github.com/LuxAeterna-git/jwt/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authorization: repo}
}
