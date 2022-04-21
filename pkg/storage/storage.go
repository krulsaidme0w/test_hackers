package pkg

import (
	"github.com/gomodule/redigo/redis"
)

func FillStorage(conn redis.Conn) {
	conn.Do("zadd", "hackers", "1953", "Richard Stallman")
	conn.Do("zadd", "hackers", "1940", "Alan Kay")
	conn.Do("zadd", "hackers", "1965", "Yukihiro Matsumoto")
	conn.Do("zadd", "hackers", "1916", "Claude Shannon")
	conn.Do("zadd", "hackers", "1969", "Linus Torvalds")
	conn.Do("zadd", "hackers", "1912", "Alan Turing")
}
