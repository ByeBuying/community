package controller

import (
	"community/protocol"
	"go-common/klay/elog"
	"net/http"

	config "community/conf"
	"community/model"

	"github.com/gin-gonic/gin"
)

type Community struct {
	ctl *Controller
	cfg *config.Config

	communityDB *model.CommunityDB
}

// NewCommunity : Community 객체 할당 및 반환.
func NewCommunity(h *Controller, rep *model.Repositories) *Community {
	r := &Community{
		ctl: h,
		cfg: h.config,
	}

	if err := rep.Get(&r.communityDB); err != nil {
		elog.Crit("newCommunity", "error", err)
	}

	return r
}

// @Router /napi/v1/community/post [post]
func (p *Community) Post(c *gin.Context) {

	req := protocol.PostWriteReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		p.ctl.RespError(c, &req, http.StatusUnprocessableEntity, "ShouldBindJSON", err)
		return
	}

}
