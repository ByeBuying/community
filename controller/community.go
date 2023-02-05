package controller

import (
	config "community/conf"
	"community/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-common/klay/elog"
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

func (p *Community) GetTest(c *gin.Context) {
	fmt.Println("aaa")
	elog.Error("error", "Test")
	elog.Info("api test", "connect")
}
