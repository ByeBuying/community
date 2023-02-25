package protocol

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
