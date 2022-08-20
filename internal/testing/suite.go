package testing

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	DB         *gorm.DB
	Redis      RedisFaker
	Service    interface{}
	Repository interface{}
}

func FakeDB() (*gorm.DB, error) {
	conn, _, _ := sqlmock.New()
	defer conn.Close()

	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  "sqlmock_pg_joatis",
		PreferSimpleProtocol: true,
		WithoutReturning:     false,
		Conn:                 conn,
	})

	db, err := gorm.Open(dialector)
	if err != nil {
		return nil, fmt.Errorf("cannot get DB object | error: %v", err)
	}

	return db, nil
}



// NewSuite for testing only
// It generates a fake db if `db` is nil
// ***** IT PANICS if cannot generate fake db *****
func NewSuite(service, repo interface{}, db *gorm.DB, redis *redis.Client) *Suite {
	var err error
	if db == nil {
		db, err = FakeDB()
		if err != nil {
			panic(err)
		}
	}

	return &Suite{
		DB:         db,
		Service:    service,
		Repository: repo,
	}
}
