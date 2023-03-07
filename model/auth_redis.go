package model

import (
	"community/conf"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-common/klay/elog"
	"strconv"
	"strings"
)

type AuthRedisDB struct {
	client *redis.Client
	cfg    *conf.Config
}

func NewAuthRedis(config *conf.Config, root *Repositories) (IRepository, error) {
	redisConfig, ok := config.Repositories["redis-auth"]["db"].(string)
	if !ok || len(redisConfig) < 0 {
		redisConfig = "0"
	}

	redisDB, err := strconv.Atoi(redisConfig)
	if err != nil {
		redisDB = 0
	}

	redisOption := redis.Options{
		Addr:     config.Repositories["redis-auth"]["datasource"].(string),
		Password: config.Repositories["redis-auth"]["pass"].(string),
		DB:       redisDB,
	}

	client := redis.NewClient(&redisOption)

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	r := &AuthRedisDB{
		client: client,
		cfg:    config,
	}

	elog.Trace("load repository : AuthRedisDB")
	return r, nil
}

func (p *AuthRedisDB) Start() error {
	return nil
}

func (p *AuthRedisDB) CheckUserClaim(key string) (*UserClaims, error) {

	redisKey := key
	result, err := p.client.Get(redisKey).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key does not exist : %v", err)
	} else if err != nil {
		// TODO: 레디스 조회 X - 반환값 수정
		panic(err)
	} else {
		// success
		resultString := strings.Split(result, ":")
		return &UserClaims{
			UserID: "user",
			Email:  resultString[1],
			Role:   resultString[3],
		}, nil
	}

}
