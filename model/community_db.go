package model

import (
	"community/conf"
	"context"
	"go-common/klay/elog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommunityDB struct {
	client *mongo.Client

	collectionPostInfo    *mongo.Collection
	collectionFileInfo    *mongo.Collection
	collectionCommentInfo *mongo.Collection
	start                 chan struct{}
}

func NewCommunityDB(config *conf.Config, root *Repositories) (IRepository, error) {
	cfg := config.Repositories["community-db"]
	r := &CommunityDB{
		start: make(chan struct{}),
	}

	credential := options.Credential{
		Username: cfg["username"].(string),
		Password: cfg["pass"].(string),
	}

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg["datasource"].(string)).SetAuth(credential)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(cfg["db"].(string))
		r.collectionPostInfo = db.Collection("post_info")
		r.collectionFileInfo = db.Collection("file_info")
		r.collectionCommentInfo = db.Collection("comment_info")
	}

	elog.Trace("load repository : CommunityDB")
	return r, nil
}

func (p *CommunityDB) Start() error {
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
