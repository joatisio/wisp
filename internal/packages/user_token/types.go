package user_token

import "github.com/joatisio/wisp/internal/models"

const TokenIDKey = "tokenId"

// Repository
// All Get methods SHOULD return nil,nil if no token/error exist
type Repository interface {
	Create(userId models.ID, t models.Token) (*models.Token, error)
	GetAllByUserId(userId models.ID) ([]models.Token, error)
	GetById(userId models.ID, tokenId models.ID) (*models.Token, error)
	GetByAccess(userId models.ID, token string) (*models.Token, error)
	GetByRefresh(userId models.ID, token string) (*models.Token, error)
	BlockById(userId models.ID, tokenId models.ID) error
	UnblockById(userId models.ID, tokenId models.ID) error
}
