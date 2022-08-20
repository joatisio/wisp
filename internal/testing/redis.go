package testing

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type RedisFaker struct {
	db   *redis.Client
	mock *redismock.ClientMock
}

func FakeRedis() (*redis.Client, error) {
	
}
