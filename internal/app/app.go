package app

import (
	"github.com/joatisio/wisp/internal/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Config   *config.Config
	Database *gorm.DB
	Logger   *zap.Logger
	Server   *Server
}

func NewApp(c *config.Config, db *gorm.DB, logger *zap.Logger, server *Server) *App {
	return &App{
		Config:   c,
		Database: db,
		Logger:   logger,
		Server:   server,
	}
}
