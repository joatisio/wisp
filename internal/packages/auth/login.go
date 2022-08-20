package auth

import (
	"errors"

	"github.com/google/uuid"
	"github.com/joatisio/wisp/internal/encryption"
	"github.com/joatisio/wisp/internal/models"
	"go.uber.org/zap"
)

func (s *Service) Login(req LoginRequest) (*LoginResponse, error) {
	u, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		s.logger.Error("cannot get user by email", zap.String("email", req.Email))

		return nil, err
	}

	if uuid.UUID(u.ID) == uuid.Nil {
		s.logger.Error("uuid expected, nil given")

		return nil, errors.New("user id cannot be nil")
	}

	if err := encryption.CheckPassword(u.Password, req.Password); err != nil {
		s.logger.Error("wrong password", zap.String("email", req.Email))

		return nil, errors.New("wrong password")
	}

	tp, err := s.jwt.GenerateTokenPair(u.ID, u.Role)
	if err != nil {
		return nil, err
	}

	_, err = s.userTokenRepo.Create(u.ID, models.Token{
		UserId:  *u,
		Access:  tp.Access,
		Refresh: tp.Refresh,
		Blocked: 0,
	})

	if err != nil {
		s.logger.Error("create token failed", zap.String("userId", u.ID.String()))

		return nil, err
	}

	return &LoginResponse{
		User:      *u,
		TokenPair: *tp,
	}, nil
}
