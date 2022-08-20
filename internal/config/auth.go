package config

import (
	"time"
)

type Auth struct {
	JWTSignKey      string
	AccessDuration  time.Duration
	RefreshDuration time.Duration
}
