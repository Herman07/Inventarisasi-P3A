package services

import (
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/repository"
	"Inventarisasi-P3A/request"
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repositoris repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: repositoris,
	}
}
func (s *AuthService) Login(request request.LoginRequest) (*models.Users, error) {
	data, err := s.authRepository.Login(request)
	if err != nil{
		return nil, err
	} else {
	}
	return data, err
}

