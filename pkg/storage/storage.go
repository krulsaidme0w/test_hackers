package pkg

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func FillStorage(client *redis.Client, ctx *context.Context) {
	client.Do(*ctx, "zadd", "hackers", "1953", "Richard Stallman")
	client.Do(*ctx, "zadd", "hackers", "1940", "Alan Kay")
	client.Do(*ctx, "zadd", "hackers", "1965", "Yukihiro Matsumoto")
	client.Do(*ctx, "zadd", "hackers", "1916", "Claude Shannon")
	client.Do(*ctx, "zadd", "hackers", "1969", "Linus Torvalds")
	client.Do(*ctx, "zadd", "hackers", "1912", "Alan Turing")
}
