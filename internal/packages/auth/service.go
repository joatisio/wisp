package auth

import (
	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/joatisio/wisp/internal/config"
	"github.com/joatisio/wisp/internal/packages/user"
	userToken "github.com/joatisio/wisp/internal/packages/user_token"
)

type Auth interface {
	// Login accepts a LoginRequest and returns Jwt and User if everything is okay
	Login(req LoginRequest) (*LoginResponse, error)

	//Logout invalidates a token
	// This token shouldn't work anymore in further authentications
	Logout(req LogoutRequest) error

	// Refresh accepts a token and generates a fresh TokenPair
	Refresh(refreshToken string) (*TokenPair, error)

	// Register creates a user and returns the created user
	Register(req RegisterRequest) (*RegisterResponse, error)
}

type Service struct {
	config        *config.Auth
	jwt           *JWT
	db            *gorm.DB
	userRepo      user.Repository
	userTokenRepo userToken.Repository
	logger        *zap.Logger
}

// TODO FIXME *Service should be Auth after implementing all methods
func NewService(c *config.Auth, jwt *JWT, db *gorm.DB, userRepo user.Repository, uTokenRepo userToken.Repository, logger *zap.Logger) *Service {
	return &Service{
		config:        c,
		jwt:           jwt,
		db:            db,
		userRepo:      userRepo,
		userTokenRepo: uTokenRepo,
		logger:        logger,
	}
}
