package router

import (
	"fmt"

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
	friendControl    *controller.Friend
}

func NewRouter(config *config.Config, ctl *controller.Controller) (*Router, error) {
	r := &Router{
		config:           config,
		communityControl: ctl.GetCommunityHandler(),
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
		fmt.Println(friend)
		friend.GET("/ok", p.friendControl.GetTest)
		friend.POST("post", p.friendControl.CreatePost)
	} // api path
	e.GET("/test", p.communityControl.GetTest)

	return e
}
