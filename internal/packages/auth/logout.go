package auth

import (
	"go.uber.org/zap"

	"github.com/joatisio/wisp/internal/packages/user"
	token "github.com/joatisio/wisp/internal/packages/user_token"
)

func (s *Service) Logout(req LogoutRequest) error {
	tok, err := s.userTokenRepo.GetByAccess(req.UserID, req.AccessToken)

	if err != nil {
		s.logger.Error("cannot get token", zap.String(user.UserIDKey, req.UserID.String()))

		return err
	}

	if err := s.userTokenRepo.BlockById(req.UserID, tok.ID); err != nil {
		s.logger.Error(
			"cannot block token",
			zap.String(user.UserIDKey, req.UserID.String()),
			zap.String(token.TokenIDKey, tok.ID.String()),
			zap.String("error", err.Error()),
		)

		return err
	}

	return nil
}
