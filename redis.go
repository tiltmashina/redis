package redis

import (
	"github.com/go-redis/redis"
)

type RedisTemplate struct {
	Client *redis.Client
}

func NewRedisClient(url string, password string, db int) *RedisTemplate {
	return &RedisTemplate{
		Client: redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
			DB:       db,
		}),
	}
}

func (s *RedisTemplate) Set(key string, value string) error {
	err := s.Client.Set(key, value, 0).Err()
	return err
}

func (s *RedisTemplate) Get(key string) (string, error) {
	val, err := s.Client.Get(key).Result()
	return val, err
}

func (s *RedisTemplate) Delete(key string) error {
	err := s.Client.Del(key).Err()
	return err
}

//ping function
func (c *RedisTemplate) Ping() (string, error) {
	val, err := c.Client.Ping().Result()
	return val, err
}
