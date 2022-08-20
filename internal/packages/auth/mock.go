package auth

import (
	"log"
	"time"

	"go.uber.org/zap"

	"github.com/joatisio/wisp/internal/config"
	"github.com/joatisio/wisp/internal/packages/user"
	"github.com/joatisio/wisp/internal/testing"
)

func setupMock(userRepo user.Repository) *testing.Suite {
	c := &config.Auth{
		JWTSignKey:      "secureKey",
		AccessDuration:  time.Duration(15 * time.Minute),
		RefreshDuration: time.Duration(14 * 24 * time.Hour),
	}

	jwt := NewJWT(c.JWTSignKey, c.AccessDuration, c.RefreshDuration)

	db, err := testing.FakeDB()
	if err != nil {
		log.Panicf("cannot get DB object | error: %v", err)
	}

	l, _ := zap.NewDevelopment(zap.Development())

	svc := NewService(c, jwt, db, userRepo, l)

	// Auth doesn't have any Repository, so a nil is passed to NewSuite
	return testing.NewSuite(svc, nil, db)
}
