package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joatisio/wisp/internal/models"
	"github.com/joatisio/wisp/internal/packages/user"
)

func TestJWT_GenerateTokenPair(t *testing.T) {
	_jwt := NewJWT("testSignKey", 15*time.Minute, 15*24*time.Hour)

	uf := user.Faker{}

	// generate 2 fake users
	fakeUsers := uf.GenerateTestUserBulk(2)

	// make id nil to force the function raise an error
	fakeUsers[1].ID = models.ID(uuid.Nil)

	tests := []struct {
		name string
		jwt  *JWT
		user *models.User
		isOk bool
	}{
		{name: "valid_token", jwt: _jwt, user: fakeUsers[0], isOk: true},
		{name: "invalid_userId", jwt: _jwt, user: fakeUsers[1], isOk: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.jwt.GenerateTokenPair(tt.user.ID, tt.user.Role)
			if err != nil && tt.isOk {
				t.Errorf("JWT.GenerateTokenPair() error = %v, isOk %v", err, tt.isOk)
				return
			}

			if got != nil {
				claims := jwt.MapClaims{}
				parsed, err := jwt.ParseWithClaims(got.Access, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(tt.jwt.SignKey), nil
				})

				if err != nil {
					t.Errorf("cannot parse access token. Error = %v", err)
				}

				if !parsed.Valid {
					t.Error("invalid access token")
				}
			}
		})
	}
}
