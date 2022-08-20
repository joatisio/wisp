package user

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/joatisio/wisp/internal/encryption"
	"github.com/joatisio/wisp/internal/models"
)

// Faker is for testing
type Faker struct {
	Password string
}

// generateTestUser for testing purposes
func (f *Faker) GenerateTestUserBulk(count int) []*models.User {
	var results []*models.User
	for i := 0; i < count; i++ {
		results = append(results, f.GenerateFakeUser(fmt.Sprintf("user%d@domain.tld", i)))
	}

	return results
}

func (f *Faker) GenerateFakeUser(email string) *models.User {
	em, err := mail.ParseAddress(email)
	if err != nil {
		panic("valid email expected")
	}

	hashPass := encryption.GetPasswordHash(f.Password)

	return &models.User{
		BaseModel: models.BaseModel{
			ID:        models.ID(uuid.New()),
			CreatedAt: time.Now(),
		},
		Username: em.Name,
		Password: hashPass,
		Email:    em.Address,
		Role:     models.RoleAdmin,
		Contact: models.Person{
			BaseModel: models.BaseModel{
				ID:        models.ID(uuid.New()),
				CreatedAt: time.Now(),
			},
			FirstName: fmt.Sprintf("%sFirstName", em.Name),
			LastName:  fmt.Sprintf("%sLastName", em.Name),
			Address:   fmt.Sprintf("%sAddress", em.Name),
			Avatar:    models.Media{},
			OwnerID:   models.ID{},
			OwnerType: "",
		},
	}
}

type MockRepo struct {
	Faker Faker
}

func NewMockRepo(f Faker) *MockRepo {
	return &MockRepo{f}
}

func MockUserRepository(f Faker) Repository {
	return NewMockRepo(f)
}

func (m *MockRepo) GetByEmail(email string) (*models.User, error) {
	return m.Faker.GenerateFakeUser(email), nil
}

func (m *MockRepo) Create(u *models.User) (*models.User, error) {
	return m.Faker.GenerateFakeUser(u.Email), nil
}

func (m *MockRepo) Update(userId models.ID, u *models.User) error {
	return nil
}

func (m *MockRepo) UpdatePassword(userId models.ID, newPassword string) error {
	return nil
}

func (m *MockRepo) Activate(userId models.ID) error {
	return nil
}

func (m *MockRepo) Deactivate(userId models.ID) error {
	return nil
}
