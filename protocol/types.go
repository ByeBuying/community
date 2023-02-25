package protocol

import (
	"strings"
	"time"
)

// 모든 응답의 헤더
type RespHeader struct {
	Result       ResultCode `json:"result"`
	ResultString string     `json:"resultString"`
	Desc         string     `json:"desc"`
}

// RespHeader : RespHeader 객체 생성 및 반환
func NewRespHeader(resultCode ResultCode, desc ...string) *RespHeader {
	return &RespHeader{
		Result:       resultCode,
		ResultString: resultCode.toString(),
		Desc:         strings.Join(desc, ","),
	}
}

// post info schema
type PostInfo struct {
	UserId    string    `json:"userId" bson:"user_id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Likes     int       `json:"likes" bson:"likes"`
	CreateAt  time.Time `json:"createAt" bson:"create_at"`
	UpdateAt  time.Time `json:"updateAt" bson:"update_at"`
	CommentId []int     `json:"commentId" bson:"comment_id"`
	FileId    int       `json:"fileId" bson:"file_id"`
	Stat      int       `json:"stat" bson:"stat"`
}
