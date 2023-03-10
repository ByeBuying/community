package model

import (
	"community/conf"
	"context"
	"go-common/klay/elog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDB struct {
	client *mongo.Client

	collectionUserInfo *mongo.Collection
	start              chan struct{}
}

func NewUserDB(config *conf.Config, root *Repositories) (IRepository, error) {
	cfg := config.Repositories["user-db"]
	r := &UserDB{
		start: make(chan struct{}),
	}

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg["datasource"].(string))); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(cfg["db"].(string))
		r.collectionUserInfo = db.Collection("user_info")
	}

	elog.Trace("load repository : UserDB")
	return r, nil
}

func (p *UserDB) Start() error {
	return func() (err error) {
		defer func() {
			if v := recover(); v != nil {
				err = v.(error)
			}
		}()
		close(p.start)
		return
	}()
}

func (p *UserDB) GetUserInfo(result *UserClaims) error {
	filter := bson.M{
		"email": result.Email,
	}

	if err := p.collectionUserInfo.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return err
	} else {
		return nil
	}
}

func (p *UserDB) CreateUserInfo(result *UserClaims) error {
	user := &UserClaims{
		UserID:  "user",
		ActType: "act",
		Email:   result.Email,
		Role:    result.Role,
		//TODO: sid 로직 추가
	}

	_, err := p.collectionUserInfo.InsertOne(context.Background(), user)
	if err != nil {
		return err
	} else {
		return nil
	}
}
