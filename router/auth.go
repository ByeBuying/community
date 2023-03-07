package router

import (
	"github.com/gin-gonic/gin"
)

// test@test.com:role:admin
// email:test@test.com:role:admin

func (p *Router) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}

	}
}

// CheckUser : 유저의 session check
func (p *Router) CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}

		// session 체크
		claim, err := p.accountControl.CheckUserSession(c)
		if err != nil {
			// 인증 에러 log 찍기
			return
		} else {
			// context setting
			c.Set("claim", claim)
			c.Next()
		}

	}

}
