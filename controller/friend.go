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

type Friend struct {
	ctl *Controller
	cfg *config.Config

	communityDB *model.CommunityDB
}

// NewCommunity : Community 객체 할당 및 반환.
func NewFriend(h *Controller, rep *model.Repositories) *Friend {
	r := &Friend{
		ctl: h,
		cfg: h.config,
	}

	if err := rep.Get(&r.communityDB); err != nil {
		elog.Crit("newCommunity", "error", err)
	}

	return r
}

// GetFriendPostList
// @Summary 친구찾기 전체 게시물 조회
// @Tags friend
// @Accept json
// @Produce json
// @Success 200 {object} protocol.FriendPostListResp
// @Router /v1/friend/list [get]
func (r *Friend) GetFriendPost(c *gin.Context) {
	var friendPostList []protocol.FriendPost

	err := r.communityDB.GetFriendPostList(&friendPostList)
	if err != nil {
		r.ctl.RespError(c, nil, http.StatusNotFound, err)
		return
	}
	// 성공 응답
	r.ctl.RespOK(c, &protocol.FriendPostListResp{
		RespHeader:     protocol.NewRespHeader(protocol.Success),
		FriendPostList: friendPostList,
	})
}

// CreateFriendPost
// @Summary 친구찾기 게시물 등록
// @Tags friend
// @Accept json
// @Produce json
// @Param author formData string true "author"
// @Param description formData string true "description"
// @Param file formData file true "file"
// @Router /v1/friend [post]
func (p *Friend) CreateFriendPost(c *gin.Context) {
	// shouldbind로 묶어볼 수 있으면 묶기
	// @Param formData formData protocol.PostReq true "Body with file "
	author := c.PostForm("author")
	description := c.PostForm("description")
	image, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	postInfo := protocol.PostReq{Author: author, Description: description, ImageName: image.Filename}
	dbErr := p.communityDB.CreateFriendPost(postInfo)
	if dbErr != nil {
		p.ctl.RespError(c, nil, http.StatusInternalServerError, "CreateFriendPost", err.Error())
		return
	}

	p.ctl.RespOK(c, &protocol.FriendPostCreateRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})

	// s3 := aws.S3Info{AwsS3Region: "", AwsAccessKey: "", AwsSecretKey: "", BucketName: ""}
	// errs := s3.SetS3ConfigByKey()
	// if errs != nil {
	// }
	// file, err := image.Open()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// s3.UploadFile(file, image.Filename, "cats")
}

// UpdateFriendPostInfo
// @Summary 친구찾기 게시물 수정
// @Tags friend
// @Accept json
// @Produce json
// @Param requestBody body protocol.PostReq true "resposne body"
// @Param id path string true "post id"
// @Router /v1/friend/{id} [put]
func (p *Friend) UpdateFriendPost(c *gin.Context) {
	id := c.Param("id")

	req := protocol.PostReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

	// db update
	err := p.communityDB.UpdateFriendPostOneById(id, req)
	if err != nil {
		p.ctl.RespError(c, nil, http.StatusInternalServerError, "UpdateReviewPost", err.Error())
		return
	}

	p.ctl.RespOK(c, &protocol.ReviewPostUpdateRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})
}

// DeleteFriendPostInfo
// @Summary 친구찾기 게시물 소프트삭제
// @Tags friend
// @Accept json
// @Produce json
// @Param id path string true "post id"
// @Router /v1/friend/{id} [delete]
func (p *Friend) DeleteFriendPost(c *gin.Context) {
	id := c.Param("id")
	err := p.communityDB.DeleteFriendPostOneById(id)
	if err != nil {
		p.ctl.RespError(c, nil, http.StatusInternalServerError, "Delete FriendPost", err.Error())
		return
	}

	p.ctl.RespOK(c, &protocol.ReviewPostUpdateRes{
		RespHeader: protocol.NewRespHeader(protocol.Success),
		Stat:       1,
	})
}
