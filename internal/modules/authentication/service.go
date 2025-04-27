package authentication

import (
	"errors"

	"SparFortuneDDD/pkg"
)

type Service struct {
	jwtService *pkg.JWTService
}

func NewService(jwtService *pkg.JWTService) *Service {
	return &Service{jwtService: jwtService}
}

var ErrInvalidUserID = errors.New("invalid user ID")

func (s *Service) GenerateToken(userID string) (string, error) {
	if userID == "" {
		return "", ErrInvalidUserID
	}
	return s.jwtService.GenerateToken(userID)
}
