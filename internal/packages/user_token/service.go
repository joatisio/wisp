package user_token

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/joatisio/wisp/internal/models"
)

type UserToken interface {
	Create(uid models.ID, token models.Token) error
	GetByUserId(uid models.ID) ([]*models.Token, error)
	Delete(tokenId models.ID) error
}

type Service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, l *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: l,
	}
}

var redisUserTokensStyle = "user:%s:tokens"
var redisTokenStyle = "token:%s"
var errCannotFindToken = errors.New("cannot find a valid token")

type RedisRepo struct {
	redis *redis.Client
	ctx   context.Context
}

func NewRedisRepo(r *redis.Client) Repository {
	return &RedisRepo{
		redis: r,
		ctx:   context.Background(),
	}
}

func (r *RedisRepo) Create(userId models.ID, t models.Token) (*models.Token, error) {
	t.ID = models.ID(uuid.New())

	tJson, err := t.Json()
	if err != nil {
		return nil, err
	}

	_, err = r.redis.HSet(
		r.ctx,
		generateUserTokensKey(userId.String()),
		generateTokenKey(t.ID.String()),
		tJson).Result()

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *RedisRepo) GetAllByUserId(userId models.ID) ([]models.Token, error) {
	return getAllByUserId(r.ctx, r.redis, userId)
}

func (r *RedisRepo) GetById(userId, tokenId models.ID) (*models.Token, error) {
	return getById(r.ctx, r.redis, userId, tokenId)
}

func (r *RedisRepo) GetByAccess(userId models.ID, token string) (*models.Token, error) {
	res, err := r.redis.HVals(
		r.ctx,
		generateUserTokensKey(userId.String())).Result()

	if err != nil {
		return nil, err
	}

	for _, j := range res {
		t, err := models.JsonToToken(j)
		if err != nil {
			return nil, err
		}

		if t.Access == token {
			return t, nil
		}
	}

	return nil, nil
}

func (r *RedisRepo) GetByRefresh(userId models.ID, token string) (*models.Token, error) {
	res, err := r.redis.HVals(
		r.ctx,
		generateUserTokensKey(userId.String())).Result()

	if err != nil {
		return nil, err
	}

	for _, j := range res {
		t, err := models.JsonToToken(j)
		if err != nil {
			return nil, err
		}

		if t.Refresh == token {
			return t, nil
		}
	}

	return nil, nil
}

func (r *RedisRepo) BlockById(userId, tokenId models.ID) error {
	return changeBlockVal(r.ctx, r.redis, userId, tokenId, 1)
}

func (r *RedisRepo) UnblockById(userId, tokenId models.ID) error {
	return changeBlockVal(r.ctx, r.redis, userId, tokenId, 0)
}

func changeBlockVal(ctx context.Context, rdb *redis.Client, userId, tokenId models.ID, val int) error {
	t, err := getById(ctx, rdb, userId, tokenId)
	if err != nil {
		return err
	}

	if t == nil {
		return errCannotFindToken
	}

	// leave it unchanged
	if t.Blocked == val {
		return nil
	}

	t.Blocked = val

	j, err := t.Json()
	if err != nil {
		return err
	}

	if _, err = rdb.HSet(
		ctx,
		generateUserTokensKey(userId.String()),
		generateTokenKey(tokenId.String()),
		j).Result(); err != nil {
		return err
	}

	return nil
}

func getById(ctx context.Context, rdb *redis.Client, userId, tokenId models.ID) (*models.Token, error) {
	res, err := rdb.HGet(
		ctx,
		generateUserTokensKey(userId.String()),
		generateTokenKey(tokenId.String())).Result()

	if err != nil {
		return nil, err
	}

	t, err := models.JsonToToken(res)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getAllByUserId(ctx context.Context, rdb *redis.Client, userId models.ID) ([]models.Token, error) {
	var tokens []models.Token

	res, err := rdb.HVals(
		ctx,
		generateUserTokensKey(userId.String())).Result()

	if err != nil {
		return nil, err
	}

	for _, tJson := range res {
		if t, err := models.JsonToToken(tJson); err != nil {
			return nil, err
		} else {
			tokens = append(tokens, *t)
		}
	}

	return tokens, nil
}

func generateUserTokensKey(userId string) string {
	return fmt.Sprintf(redisUserTokensStyle, userId)
}

func generateTokenKey(tokenId string) string {
	return fmt.Sprintf(redisTokenStyle, tokenId)
}
