package redis

import (
	"encoding/json"
	"fmt"
	"goutils/errors"
	"time"

	"github.com/go-redis/redis/v7"
)

type Client interface {
	Set(key string, value interface{}, ttl time.Duration) error
	SetStruct(key string, value interface{}, ttl time.Duration) error
	GetBytes(key string) ([]byte, error)
	GetString(key string) (string, error)
	GetStruct(key string, v interface{}) error
	Del(key string) error
	Close() error
	Client() *redis.Client
}

type client struct {
	redis *redis.Client
}

func NewClient(conf Config) (Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         conf.Addr(),
		DialTimeout:  time.Duration(conf.ConnectionTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(conf.ConnectionTimeout) * time.Millisecond,
		MinIdleConns: conf.MaxIdleConnections,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &client{
		redis: redisClient,
	}, nil
}

func (c *client) Set(key string, value interface{}, ttl time.Duration) error {
	return c.redis.Set(key, value, ttl).Err()
}

func (c *client) SetStruct(key string, value interface{}, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return c.redis.Set(key, jsonData, ttl).Err()
}

func (c *client) GetBytes(key string) ([]byte, error) {
	value, err := c.redis.Get(key).Bytes()

	if err != nil {
		return nil, cacheError(key, err)
	}

	return value, nil
}

func (c *client) GetString(key string) (string, error) {
	value, err := c.redis.Get(key).Result()

	if err != nil {
		return "", cacheError(key, err)
	}

	return value, nil
}

func (c *client) GetStruct(key string, v interface{}) error {
	value, err := c.GetBytes(key)

	if err != nil {
		return err
	}

	return json.Unmarshal(value, v)
}

func (c *client) Close() error {
	return c.redis.Close()
}

func cacheError(key string, err error) error {
	if err == redis.Nil {
		return errors.NewError("cache_miss", fmt.Sprintf("key '%s' does not exist", key), nil, err)
	}
	return err
}

func (c *client) Client() *redis.Client {
	return c.redis
}

func (c *client) Del(key string) error {
	return c.redis.Del(key).Err()
}
