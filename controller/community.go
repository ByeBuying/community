package controller

import (
	"fmt"
	"go-common/klay/elog"

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

// Get
// @Summary Test
// @Description Test
// @Accept  json
// @Produce  json
// @Router /test [get]
func (p *Community) GetTest(c *gin.Context) {
	fmt.Println("aaa")
	elog.Error("error", "Test")
	elog.Info("api test", "connect")
}
