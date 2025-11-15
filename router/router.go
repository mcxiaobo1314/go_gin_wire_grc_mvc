package router

import (
	"gin_gorm/controllers"
	"gin_gorm/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterProviderSet = wire.NewSet(NewRouter)

func NewRouter(Auth *controllers.AuthController, DemoMiddleware *middlewares.DemoMiddleware) *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth").Use(DemoMiddleware.Ac())
	{
		auth.GET("/test", Auth.Login)
		auth.GET("/testv2/:id", Auth.Loginv2)
	}

	return r

}
