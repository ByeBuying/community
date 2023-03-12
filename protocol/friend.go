package protocol

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendPost struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    string             `json:"userId" bson:"user_id"`
	Content   string             `json:"content" bson:"content"`
	ImageUrl  string             `json:"imageUrl" bson:"image_url"`
	Likes     int                `json:"likes" bson:"likes"`
	LikeUsers []string           `json:"likeUsers" bson:"like_users"`
	Comments  []string           `json:"comments" bson:"comments`
	CreateAt  time.Time          `json:"createAt" bson:"create_at"`
	UpdateAt  time.Time          `json:"updateAt" bson:"update_at"`
	Stat      int                `json:"stat" bson:"stat"`
}

type FriendPostAndComments struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	UserId        string             `json:"userId" bson:"user_id"`
	Content       string             `json:"content" bson:"content"`
	ImageUrl      string             `json:"imageUrl" bson:"image_url"`
	Likes         int                `json:"likes" bson:"likes"`
	LikeUsers     []string           `json:"likeUsers" bson:"like_users"`
	Comments      []string           `json:"comments" bson:"comments`
	FriendComment []FriendComment    `json:"friendComment" bson:"friend_comment"`
	CreateAt      time.Time          `json:"createAt" bson:"create_at"`
	UpdateAt      time.Time          `json:"updateAt" bson:"update_at"`
	Stat          int                `json:"stat" bson:"stat"`
}

type FriendComment struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       string             `json:"userId" bson:"user_id"`
	PostSelector string             `json:"postSelector" bson:"post_selector"`
	PostId       primitive.ObjectID `json:"postId" bson:"post_id"`
	Content      string             `json:"content" bson:"content"`
	CreateAt     time.Time          `json:"createAt" bson:"create_at"`
	UpdateAt     time.Time          `json:"updateAt" bson:"update_at"`
	Stat         int                `json:"stat" bson:"stat"`
}

// todo 나중에 스키마랑 res / req struct 구분
// ---- 스키마 구분  ----

type PostReq struct {
	Author      string
	Description string
	ImageName   string
}

type FriendCommentReq struct {
	UserId  string `json:"userId" bson:"user_id"`
	Content string `json:"content" bson:"content"`
}
type FriendPostListResp struct {
	*RespHeader
	FriendPostList []FriendPostAndComments `json:"friendPostList"`
}

type FriendPostCreateRes struct {
	*RespHeader
	Stat int `json:"stat"`
}
