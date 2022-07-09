package app

import (
	"github.com/joatisio/wisp/internal/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	SetDB(db *gorm.DB)
}

type API interface {
}

type App struct {
	Config   *config.Config
	Database *gorm.DB
	Logger   *zap.Logger
}

func NewApp(c *config.Config, db *gorm.DB, logger *zap.Logger) *App {
	return &App{
		Config:   c,
		Database: db,
		Logger:   logger,
	}
}
