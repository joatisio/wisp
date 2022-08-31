package user

import (
	"github.com/joatisio/wisp/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User interface {
	Create(user models.User) error
	Get(id models.ID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(id models.ID, user models.User) error
	Delete(id models.ID) error
}

type Service struct {
	repo   Repository
	db     *gorm.DB
	logger *zap.Logger
}
