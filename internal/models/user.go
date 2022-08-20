package models

import "strings"

type User struct {
	BaseModel

	Username string `json:"username"`
	Password string `json:"-"`

	Email string

	Contact Person `gorm:"polymorphic:Owner;"`

	Role UserRole

	ConfirmedAt NullTime
}

type UserRole string

const (
	RoleOwner    UserRole = "owner"
	RoleAdmin    UserRole = "admin"
	RoleCreator  UserRole = "creator"
	RoleObserver UserRole = "observer"
)

func (r UserRole) String() string {
	return string(r)
}

func (r UserRole) IsValid() bool {
	switch r {
	case RoleOwner:
		return true
	case RoleAdmin:
		return true
	case RoleCreator:
		return true
	case RoleObserver:
		return true
	}

	return false
}

func (r *UserRole) Clean() {
	if r == nil {
		// Just in case
		return
	}

	s := r.String()
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	*r = UserRole(s)
}
