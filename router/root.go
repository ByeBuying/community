package router

import (
	config "community/conf"
	"community/controller"
	"community/util/recovery"

	"community/docs"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	config *config.Config

	communityControl *controller.Community

	friendControl *controller.Friend

	reviewControl *controller.Review
}

func NewRouter(config *config.Config, ctl *controller.Controller) (*Router, error) {
	r := &Router{
		config:           config,
		communityControl: ctl.GetCommunityHandler(),
		reviewControl:    ctl.GetReviewHandler(),
		friendControl:    ctl.GetFriendHandler(),
	}

	return r, nil
}

// func CORS() gin.HandlerFunc {
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {
	// e := gin.New()
	if p.config.Common.ServiceId == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(recovery.GinRecovery(p.config.Common.ServiceId))
	e.Use(CORS())

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Title = "community"
	url := ginSwg.URL("http://localhost:8080/swagger/doc.json")
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler, url))

	friend := e.Group("v1/friend/")
	{
		friend.GET("/list", p.friendControl.GetFriendPost)
		friend.POST("/", p.friendControl.CreateFriendPost)
		friend.PUT("/:id", p.friendControl.UpdateFriendPost)
		friend.DELETE("/:id", p.friendControl.DeleteFriendPost)
	}

	community := e.Group("/napi/v1/community")
	{
		community.POST("/post", p.communityControl.Post)
	}

	// review
	review := e.Group("/review/v1/post")
	{
		// TODO middleware 추가
		review.GET("/list", p.reviewControl.GetPostList)
		review.GET("/:id", p.reviewControl.GetPostDetail)
		review.POST("", p.reviewControl.CreatePostInfo)
		review.PUT("/:id", p.reviewControl.UpdatePostInfo)
		review.PATCH("/:id", p.reviewControl.DeletePostInfo)
		review.POST("/like/:id", p.reviewControl.ChangeLike)
	}

	return e
}
