package protocol

import "time"

type PostWriteReq struct {
	UserId   string    `json:"userId"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time `json:"deleteAt"`
	Stat     int       `json:"stat"`
}

type PostWriteRes struct {
	*RespHeader
	Stat int `json:"stat"`
}
