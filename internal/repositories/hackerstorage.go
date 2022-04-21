package repositories

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
	"test_hackers/internal/core/domain"
)

type Storage struct {
	conn redis.Conn
}

func NewStorage(conn redis.Conn) *Storage {
	return &Storage{
		conn: conn,
	}
}

func (s Storage) GetAll() ([]domain.Hacker, error) {
	resp, err := redis.Strings(s.conn.Do("ZRANGEBYSCORE", "hackers", "-INF", "+INF", "WITHSCORES"))
	if err != nil {
		return []domain.Hacker{}, err
	}

	hackers := make([]domain.Hacker, 0, len(resp)/2)
	for i := 0; i < len(resp); i = i + 2 {
		name := resp[i]
		score, err := strconv.Atoi(resp[i+1])
		if err != nil {
			return []domain.Hacker{}, err
		}

		hackers = append(hackers, domain.Hacker{
			Name:  name,
			Score: score,
		})
	}

	return hackers, nil
}
