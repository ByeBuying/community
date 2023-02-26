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
	CreateAt  time.Time          `json:"createAt" bson:"create_at"`
	UpdateAt  time.Time          `json:"updateAt" bson:"update_at"`
	Stat      int                `json:"stat" bson:"stat"`
}

type FriendPostListResp struct {
	*RespHeader
	FriendPostList []FriendPost `json:"friendPostList"`
}
