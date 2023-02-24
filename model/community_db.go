package model

import (
	"community/conf"
	"community/protocol"
	"context"
	"go-common/klay/elog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CommunityDB struct {
	client *mongo.Client

	collectionPostInfo    *mongo.Collection
	collectionFileInfo    *mongo.Collection
	collectionCommentInfo *mongo.Collection
	collectionReviewInfo  *mongo.Collection
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
		r.collectionReviewInfo = db.Collection("review_info")
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

func (p *CommunityDB) CreatePost(req protocol.PostWriteReq) error {
	//result := protocol.PostInfo{}
	//_, err := p.collectionPostInfo.InsertOne(context.Background(), )
	return nil
}

func (p *CommunityDB) GetReviewList(result *[]protocol.ReviewPost) error {
	filter := bson.M{
		"stat": 1,
	}
	sort := bson.D{{"create_at", 1}}
	findOptions := options.Find().SetSort(sort)

	if cursor, err := p.collectionReviewInfo.Find(context.Background(), filter, findOptions); err != nil {
		return err
	} else {
		defer cursor.Close(context.Background())
		cursor.All(context.Background(), result)
		return nil
	}
}

func (p *CommunityDB) GetReviewDetail(id string, result *protocol.ReviewPost) error {

	convertId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":  convertId,
		"stat": 1,
	}

	err = p.collectionReviewInfo.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (p *CommunityDB) CreateReviewPost(req protocol.ReviewPostReq) error {

	post := protocol.ReviewPost{
		Id:       primitive.NewObjectID(),
		Title:    req.Title,
		UserId:   req.UserId,
		Content:  req.Content,
		Likes:    0,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Stat:     1,
		Grade:    req.Grade,
	}

	_, err := p.collectionReviewInfo.InsertOne(context.Background(), post)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (p *CommunityDB) UpdateReviewPost(id string, req protocol.ReviewPostReq) error {
	convertId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":  convertId,
		"stat": 1,
	}
	update := bson.M{
		"$set": bson.M{
			"title":     req.Title,
			"content":   req.Content,
			"update_at": time.Now(),
			"grade":     req.Grade,
		},
	}

	_, err = p.collectionReviewInfo.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	} else {
		return nil
	}

}

func (p *CommunityDB) DeleteReviewPost(id string) error {
	convertId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":  convertId,
		"stat": 1,
	}
	update := bson.M{
		"$set": bson.M{
			"stat":      0,
			"update_at": time.Now(),
		},
	}

	_, err = p.collectionReviewInfo.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	} else {
		return nil
	}
}
