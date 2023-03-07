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

	friendControl  *controller.Friend
	reviewControl  *controller.Review
	accountControl *controller.Account
}

func NewRouter(config *config.Config, ctl *controller.Controller) (*Router, error) {
	r := &Router{
		config:           config,
		communityControl: ctl.GetCommunityHandler(),
		reviewControl:    ctl.GetReviewHandler(),
		friendControl:    ctl.GetFriendHandler(),
		accountControl:   ctl.GetAccountHandler(),
	}

	return r, nil
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

	// e.GET("/health", p.healthControl.Check)
	// swagger
	docs.SwaggerInfo.Host = "localhost:8080" // swagger 정보 등록
	docs.SwaggerInfo.Title = "community"
	url := ginSwg.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definitionv
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler, url))

	//if p.config.Common.ServiceId == "alpha" {
	//	e.GET("/swagger/:any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
	e.GET("/image", func(c *gin.Context) {
		c.File("./public/image/test.png")
		// c.JSON(200, gin.H{"result": "ok"})
	})
	friend := e.Group("friend/v1")
	{
		friend.GET("/post/list", p.friendControl.GetFriendPost)
		friend.POST("post", p.friendControl.CreatePost)
		friend.PUT("post/:id", p.friendControl.UpdatePost)
		friend.DELETE("post/:id", p.friendControl.DeletePost)
	} // api path

	// api path

	// community
	community := e.Group("/napi/v1/community")
	{
		community.POST("/post", p.communityControl.Post)
	}

	// review
	review := e.Group("/review/v1/post")
	{
		// TODO middleware 추가
		review.GET("/list", p.CheckUser(), p.reviewControl.GetPostList)
		review.GET("/:id", p.reviewControl.GetPostDetail)
		review.POST("", p.reviewControl.CreatePostInfo)
		review.PUT("/:id", p.reviewControl.UpdatePostInfo)
		review.PATCH("/:id", p.reviewControl.DeletePostInfo)
		review.POST("/like/:id", p.reviewControl.ChangeLike)
	}

	return e
}
