package controller

import (
	"fmt"
	"go-common/klay/elog"
	"net/http"
	"os"

	config "community/conf"
	"community/model"
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

// Get
// @Summary Test
// @Description Test
// @Accept  json
// @Produce  json
// @Router /test [get]
func (p *Friend) GetTest(c *gin.Context) {
	c.JSON(200, gin.H{"result": "ok`"})
	// fmt.Println("aaa")
	// elog.Error("error", "Test")
	// elog.Info("api test", "connect")
}

func (p *Friend) CreatePost(c *gin.Context) {
	// shouldbind로 묶어볼 수 있으면 묶기
	author := c.PostForm("author")
	text := c.PostForm("text")
	// Get image
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
		"text":   text,
	})
}
