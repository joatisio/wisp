package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joatisio/wisp/internal/models"
)

var (
	errNilUserId = errors.New("userId is nil")
)

type JWT struct {
	SignKey         string
	AccessDuration  time.Duration
	RefreshDuration time.Duration
}

func NewJWT(signKey string, access, refresh time.Duration) *JWT {
	return &JWT{
		SignKey:         signKey,
		AccessDuration:  access,
		RefreshDuration: refresh,
	}
}

func (j *JWT) GenerateTokenPair(userId models.ID, role models.UserRole) (*TokenPair, error) {

	if userId.IsNil() {
		return nil, errNilUserId
	}

	var signingKey = []byte(j.SignKey)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(j.AccessDuration).Unix()
	claims["sub"] = uuid.UUID(userId).String()
	claims["role"] = role

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = userId
	rtClaims["exp"] = time.Now().Add(j.RefreshDuration).Unix()

	rt, err := refreshToken.SignedString(signingKey)
	if err != nil {
		return nil, fmt.Errorf("cannot generate access token: %s", err.Error())
	}

	access, err := token.SignedString(signingKey)
	if err != nil {
		return nil, fmt.Errorf("cannot generate refresh token: %s", err.Error())
	}

	return &TokenPair{
		Access:  access,
		Refresh: rt,
	}, nil
}
