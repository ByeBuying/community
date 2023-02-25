package protocol

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// review post info schema
type ReviewPost struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	UserId    string             `json:"userId" bson:"user_id"`
	Content   string             `json:"content" bson:"content"`
	Likes     int                `json:"likes" bson:"likes"`
	LikeUsers []string           `json:"likeUsers" bson:"like_users"`
	CreateAt  time.Time          `json:"createAt" bson:"create_at"`
	UpdateAt  time.Time          `json:"updateAt" bson:"update_at"`
	Stat      int                `json:"stat" bson:"stat"`
	Grade     float64            `json:"grade" bson:"grade"`
}

type ReviewListResp struct {
	*RespHeader
	ReviewList []ReviewPost `json:"reviewList"`
}

type ReviewDetailResp struct {
	*RespHeader
	ReviewDetail ReviewPost `json:"reviewDetail"`
}

type ReviewPostReq struct {
	Title   string  `json:"title"`
	UserId  string  `json:"userId"`
	Content string  `json:"content"`
	Stat    int     `json:"stat"`
	Grade   float64 `json:"grade"`
}

type ReviewPostCreateRes struct {
	*RespHeader
	Stat int `json:"stat"`
}

type ReviewPostUpdateRes struct {
	*RespHeader
	Stat int `json:"stat"`
}

type ReviewPostDeleteRes struct {
	*RespHeader
	Stat int `json:"stat"`
}
