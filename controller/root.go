package controller

import (
	config "community/conf"
	"community/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller : handler api 라우터.

type Controller struct {
	config *config.Config

	communityController *Community
}

func New(config *config.Config, port int, rep *model.Repositories) (*Controller, error) {
	r := &Controller{
		config: config,
	}

	// auth redis 먼저 확인
	//if err := rep.Get()

	r.communityController = NewCommunity(r, rep)

	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

// 수정 중
func (p *Controller) RespError(c *gin.Context, body interface{}, status int, err ...interface{}) {
	//bytes, _ := json.Marshal(body)
	//userClaim, _ := c.Get(Claim)
	//
	//elog.Error("Request error", "path", c.FullPath(), "body", bytes, "status", status, "error", joinMsg(err), "userClaim", userClaim)
	//sentry.CaptureException(errors.New(joinMsg(err)))
	//c.JSON(status, protocol.NewRespHeader(protocol.Failed, joinMsg(c.Request.URL, err)))
	//c.Abort()
}

// get post handler
func (p *Controller) GetCommunityHandler() *Community {
	return p.communityController
}
