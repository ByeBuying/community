package model

import (
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	UserID  string `json:"userId" bson:"user_id"`
	ActType string `json:"actType" bson:"act_type"`
	Email   string `json:"email" bson:"email"`
	Role    string `json:"role" bson:"role"`
	SID     string `json:"sid" bson:"sid"` // session_id, JWT 토큰 생성시 중복 방지를 위함 그외 사용하는곳 없음
	jwt.StandardClaims
}
