package controller

import (
	"fmt"
	"go-common/klay/elog"
	"net/http"

	config "community/conf"
	"community/model"

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
	author := c.PostForm("author")
	text := c.PostForm("text")
	fmt.Println(text == "")
	// Get image
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	// 파일 저장 경로를 지정한다.
	dst := "./" + file.Filename

	// 파일을 업로드한다.
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))

	c.JSON(200, gin.H{
		"author": author,
		"image":  "123",
		"text":   text,
	})
}
