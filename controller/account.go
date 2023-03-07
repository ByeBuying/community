package controller

import (
	config "community/conf"
	"community/model"
	"github.com/gin-gonic/gin"
	"go-common/klay/elog"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ctl *Controller
	cfg *config.Config

	authRedisDB *model.AuthRedisDB
	userDB      *model.UserDB
}

func NewAccount(h *Controller, rep *model.Repositories) *Account {
	r := &Account{
		ctl: h,
		cfg: h.config,
	}

	if err := rep.Get(&r.authRedisDB, &r.userDB); err != nil {
		elog.Crit("newAccount", "error", err)
	}
	return r
}

// redis 에서 가져온 데이터 유저 포멧 변환
func (p *Account) CheckUserSession(c *gin.Context) (*model.UserClaims, error) {
	result, err := p.authRedisDB.CheckUserClaim("auth_redis")
	if err != nil {
		return nil, err
	}

	// user DB 조회
	err = p.userDB.GetUserInfo(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// user data 추가
			p.userDB.CreateUserInfo(result)
		} else {
			return nil, err
		}
	}

	elog.Info("user info", "access", result.Email)
	return result, nil
}
