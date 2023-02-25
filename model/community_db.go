package model

import (
	"context"
	"errors"
	"fmt"
	"go-common/klay/elog"
	"time"

	"community/conf"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// credential := options.Credential{
	// 	Username: cfg["username"].(string),
	// 	Password: cfg["pass"].(string),
	// }

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg["datasource"].(string))); err != nil {
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

func (p *CommunityDB) Find() {
	filter := bson.D{{}}
	p.collectionPostInfo.Find(context.TODO(), filter)
}

func (p *CommunityDB) DeleteOneById(id string) (bool, error) {
	uuid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{
		"_id": uuid,
	}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}
	// softDelete
	res, err := p.collectionPostInfo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if res.ModifiedCount == 1 {
		return true, nil
	}
	return false, errors.New("error")
}

func (p *CommunityDB) UpdateOneById(id string, description string) (bool, error) {
	uuid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{
		"_id": uuid,
	}
	update := bson.M{
		"$set": bson.M{
			"description": description,
		},
	}
	// softDelete
	res, err := p.collectionPostInfo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if res.ModifiedCount == 1 {
		return true, nil
	}
	return false, errors.New("error")
}
