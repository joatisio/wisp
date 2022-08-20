package user

import "github.com/joatisio/wisp/internal/models"

const (
	UserIDKey = "userId"
)

type Repository interface {
	GetByEmail(email string) (*models.User, error)
	Create(u *models.User) (*models.User, error)
	Update(userId models.ID, u *models.User) error
	UpdatePassword(userId models.ID, newPassword string) error
	Activate(userId models.ID) error
	Deactivate(userId models.ID) error
}
