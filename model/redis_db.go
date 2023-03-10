package model

import (
	"go-common/klay/elog"

	config "community/conf"

	"github.com/go-redis/redis/v7"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(config *config.Config, root *Repositories) (IRepository, error) {
	redisOption := redis.Options{
		Addr:     config.Repositories["redis-db"]["datasource"].(string),
		Password: config.Repositories["redis-db"]["pass"].(string),
		DB:       0,
		// TLSConfig: &tls.Config{InsecureSkipVerify: false},
	}

	client := redis.NewClient(&redisOption)
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	r := &RedisDB{
		client: client,
	}

	elog.Trace("load repository : RedisDB")
	return r, nil
}

func (p *RedisDB) Start() error {
	return nil
}

func (p *RedisDB) GetCache(key string) (string, error) {

	user, err := p.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return user, nil
}

func (p *RedisDB) DeleteCache(key string) error {

	if err := p.client.Del(key).Err(); err != nil {
		return err
	}
	return nil
}

func (p *RedisDB) HGetMember(key string) string {
	res := p.client.HGet("member", key)
	return res.Val()
}