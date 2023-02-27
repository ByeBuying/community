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
