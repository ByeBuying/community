package controller

import (
	"fmt"
	"go-common/klay/elog"
	"net/http"

	config "community/conf"
	"community/model"
	"community/protocol"

	"github.com/gin-gonic/gin"
)

type Review struct {
	ctl *Controller
	cfg *config.Config

	// DB
	communityDB *model.CommunityDB
}

// NewReview 객체 생성
func NewReview(h *Controller, rep *model.Repositories) *Review {
	r := &Review{
		ctl: h,
		cfg: h.config,
	}

	if err := rep.Get(&r.communityDB); err != nil {
		elog.Crit("newCommunity", "error", err)
	}

	return r
}

// GetPostList
// @Summary 리뷰 전체 게시물 조회
// @Description grade - 평점
// @Tags review
// @Accept json
// @Produce json
// @Success 200 {object} protocol.ReviewListResp
// @Router /review/v1/post/list [get]
func (r *Review) GetPostList(c *gin.Context) {
	var reviewInfoList []protocol.ReviewPost

	fmt.Println(&reviewInfoList, "review")
	err := r.communityDB.GetReviewList(&reviewInfoList)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusNotFound, err)
		return
	}

	// 성공 응답
	r.ctl.RespOK(c, &protocol.ReviewListResp{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		ReviewList: reviewInfoList,
	})
}

// GetPostDetail
// @Summary 리뷰 게시물 조회
// @Description grade - 평점
// @Tags review
// @Accept json
// @Produce json
// @Success 200 {object} protocol.ReviewDetailResp
// @Router /review/v1/post/{id} [get]
func (r *Review) GetPostDetail(c *gin.Context) {
	id := c.Param("id")

	var reviewInfo protocol.ReviewPost

	err := r.communityDB.GetReviewDetail(id, &reviewInfo)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusNotFound, err)
		return
	}

	r.ctl.RespOK(c, &protocol.ReviewDetailResp{
		RespHeader:   protocol.NewRespHeader(protocol.Success),
		ReviewDetail: reviewInfo,
	})
}

// CreatePostInfo
// @Summary 리뷰 게시물 생성
// @Description grade - 평점
// @Tags review
// @Accept json
// @Produce json
// @Param requestBody body protocol.ReviewPostReq true "resposne body"
// @Success 200 {object} protocol.ReviewPostCreateRes
// @Router /review/v1/post [post]
func (r *Review) CreatePostInfo(c *gin.Context) {
	req := protocol.ReviewPostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		r.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

	// TODO 유저 레디스 값 체크

	// TODO 유효성 체크

	// db insert
	err := r.communityDB.CreateReviewPost(req)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusBadRequest, "CreateReviewPost", err.Error())
		return
	}

	r.ctl.RespOK(c, &protocol.ReviewPostCreateRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})
}

// UpdatePostInfo
// @Summary 리뷰 게시물 수정
// @Description grade - 평점
// @Tags review
// @Accept json
// @Produce json
// @Param requestBody body protocol.ReviewPostReq true "resposne body"
// @Success 200 {object} protocol.ReviewPostUpdateRes
// @Router /review/v1/post/{id} [put]
func (r *Review) UpdatePostInfo(c *gin.Context) {
	id := c.Param("id")

	req := protocol.ReviewPostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		r.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

	// TODO 유저 레디스 값 체크

	// TODO 유효성 체크

	// db update
	err := r.communityDB.UpdateReviewPost(id, req)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusBadRequest, "UpdateReviewPost", err.Error())
		return
	}

	r.ctl.RespOK(c, &protocol.ReviewPostUpdateRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})
}

// DeletePostInfo
// @Summary 리뷰 게시물 삭제
// @Description grade - 평점
// @Description soft delete
// @Tags review
// @Accept json
// @Produce json
// @Param requestBody body protocol.ReviewPostReq true "resposne body"
// @Success 200 {object} protocol.ReviewPostDeleteRes
// @Router /review/v1/post/{id} [patch]
func (r *Review) DeletePostInfo(c *gin.Context) {
	id := c.Param("id")

	req := protocol.ReviewPostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		r.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

	// TODO 유저 레디스 값 체크

	// TODO 유효성 체크

	// db delete
	err := r.communityDB.DeleteReviewPost(id)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusBadRequest, "UpdateReviewPost", err.Error())
		return
	}

	r.ctl.RespOK(c, &protocol.ReviewPostDeleteRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})
}

func (r *Review) ChangeLike(c *gin.Context) {
	// 해당 글 id find
	// 유저가 좋아요 누른 게시물 리스트에 있으면 -1
	// 없으면 +1

	id := c.Param("id")
	_ = id

	req := protocol.ReviewPostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		r.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

	// review post info 에서 likeusers 에서 해당 유저의 id find
	// 없으면 +1
	// 있으면 -1
}

func (r *Review) Comment(c *gin.Context) {
}
