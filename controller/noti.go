package controller

import (
	config "community/conf"
	"community/model"
	"go-common/klay/elog"
)

type Notice struct {
	ctl *Controller
	cfg *config.Config

	communityDB *model.CommunityDB
}

// NewNotice : Notice 객체 할당 및 반환.
func NewNotice(h *Controller, rep *model.Repositories) *Notice {
	r := &Notice{
		ctl: h,
		cfg: h.config,
	}

	if err := rep.Get(&r.communityDB); err != nil {
		elog.Crit("newCommunity", "error", err)
	}

	return r
}
