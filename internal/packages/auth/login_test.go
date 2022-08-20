package auth

import (
	"testing"

	"github.com/joatisio/wisp/internal/packages/user"
)

func TestService_Login(t *testing.T) {
	password := "securePassword"
	userRepo := user.MockUserRepository(user.Faker{Password: password})
	suite := setupMock(userRepo)
	// TODO it is wrong. IT SHOULD BE: svc, ok := suite.Service.(Auth)
	// FIXME change after all methods are implemented by MockAuth
	svc, ok := suite.Service.(*Service)
	if !ok {
		t.Error("invalid type assertion")
	}

	email := "test@domain.tld"
	l, err := svc.Login(LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		t.Fatal(err)
	}

	if l.User.Email != email {
		t.Errorf("auth:Login() failed, expected: %v got: %v", email, l.User.Email)
	}

	//if l.TokenPair
}
