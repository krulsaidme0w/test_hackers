package repositories

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"test_hackers/internal/core/domain"
)

type Storage struct {
	client *redis.Client
	ctx    *context.Context
}

func NewStorage(client *redis.Client, ctx *context.Context) *Storage {
	return &Storage{
		client: client,
		ctx:    ctx,
	}
}

func (s Storage) GetAll() ([]domain.Hacker, error) {
	vals, err := s.client.ZRangeByScoreWithScores(*s.ctx, "hackers", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		return []domain.Hacker{}, err
	}

	hackers := make([]domain.Hacker, 0, len(vals))
	for _, val := range vals {
		hackers = append(hackers, domain.Hacker{
			Name:  val.Member.(string),
			Score: val.Score,
		})
	}

	return hackers, nil
}
