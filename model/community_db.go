package model

import (
	"context"
	"fmt"
	"go-common/klay/elog"
	"time"

	"community/protocol"

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
	collectionReviewInfo  *mongo.Collection
	collectionFriendInfo  *mongo.Collection
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
		r.collectionReviewInfo = db.Collection("review_info")
		r.collectionFriendInfo = db.Collection("friend_info")

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

func (p *CommunityDB) GetFriendPostList(result *[]protocol.FriendPost) error {
	filter := bson.M{
		"stat": 1,
	}
	sort := bson.D{{"create_at", 1}}
	findOptions := options.Find().SetSort(sort)

	if cursor, err := p.collectionFriendInfo.Find(context.Background(), filter, findOptions); err != nil {
		return err
	} else {
		defer cursor.Close(context.Background())
		cursor.All(context.Background(), result)
		return nil
	}
}

func (p *CommunityDB) CreateFriendPost(req protocol.PostReq) error {
	post := protocol.FriendPost{
		Id:        primitive.NewObjectID(),
		UserId:    "1",
		Content:   req.Description,
		ImageUrl:  req.ImageName,
		Likes:     0,
		LikeUsers: []string{},
		Comments:  []string{},
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		Stat:      1,
	}
	_, err := p.collectionCommentInfo.InsertOne(context.Background(), post)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (p *CommunityDB) DeleteFriendPostOneById(id string) error {
	convertId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": convertId,
	}
	update := bson.M{
		"$set": bson.M{
			"stat":      0,
			"update_at": time.Now(),
		},
	}
	// softDelete
	_, err = p.collectionPostInfo.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (p *CommunityDB) UpdateFriendPostOneById(id string, req protocol.PostReq) error {
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
			"author":    req.Author,
			"content":   req.Description,
			"update_at": time.Now(),
		},
	}

	_, err = p.collectionPostInfo.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (p *CommunityDB) CreateComment(req protocol.FriendCommentReq) error {
	// comment를 생성해주고 저장하기
	comment := protocol.FriendComment{
		Id:           primitive.NewObjectID(),
		UserId:       "userId",
		PostSelector: "friend",
		PostId:       "640842423455da0cf1f46d72",
		Content:      req.Content,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		Stat:         1,
	}

	insertRes, err := p.collectionCommentInfo.InsertOne(context.Background(), comment)
	if err != nil {
		return err
	} else {
		convertId, err := primitive.ObjectIDFromHex(comment.PostId)
		if err != nil {
			return err
		}

		filter := bson.M{
			"_id": convertId,
		}
		update := bson.M{
			"$push": bson.M{
				"comments": insertRes.InsertedID.(primitive.ObjectID).String(),
			},
		}
		_, err = p.collectionPostInfo.UpdateOne(context.Background(), filter, update)
		if err != nil {
			fmt.Println(err, "error")
			return err
		}
		return nil
	}
	// commentId를 post(Comments)에 담아두기
}

func (p *CommunityDB) GetReviewList(result *[]protocol.ReviewPost) error {
	fmt.Println("reivew hello")
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
