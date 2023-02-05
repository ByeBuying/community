package model

import (
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	UserID  string `json:"user_id"`
	ActType string `json:"act_type"`
	Email   string `json:"email"`
	SID     string `json:"sid"` // session_id, JWT 토큰 생성시 중복 방지를 위함 그외 사용하는곳 없음
	jwt.StandardClaims
}
