package controller

import (
	"errors"
	"fmt"
	"go-common/klay/elog"
	"net/http"
	"os"

	config "community/conf"
	"community/model"
	"community/protocol"
	aws "community/util"

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
// @Router /friend/v1/post/list [get]
func (r *Friend) GetFriendPost(c *gin.Context) {
	var friendPostList []protocol.FriendPost

	fmt.Println("DB")
	err := r.communityDB.GetFriendPostList(&friendPostList)
	if err != nil {
		fmt.Println(err, "err")
		r.ctl.RespError(c, nil, http.StatusNotFound, err)
		return
	}

	// 성공 응답
	r.ctl.RespOK(c, &protocol.FriendPostListResp{
		RespHeader:     protocol.NewRespHeader(protocol.Success),
		FriendPostList: friendPostList,
	})
}

func (p *Friend) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	descripiton := c.PostForm("descripiton")
	c.JSON(200, gin.H{"result": descripiton, id: id})
}

func (p *Friend) DeletePost(c *gin.Context) {
	id := c.Param("id")
	result, err := p.communityDB.DeleteOneById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": errors.New("error")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (p *Friend) CreatePost(c *gin.Context) {
	// shouldbind로 묶어볼 수 있으면 묶기
	author := c.PostForm("author")
	descripiton := c.PostForm("text")

	// Get image -> images
	image, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// toml로 관리하기
	s3 := aws.S3Info{AwsS3Region: "", AwsAccessKey: "", AwsSecretKey: "", BucketName: ""}
	errs := s3.SetS3ConfigByKey()
	if errs != nil {
	}
	file, err := image.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	s3.UploadFile(file, image.Filename, "cats")

	c.JSON(200, gin.H{
		"author": author,
		"image":  "123",
		"text":   descripiton,
	})
}
