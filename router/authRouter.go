package router

import (
	"gin_gorm/controllers"
	"gin_gorm/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var AuthRouterProviderSet = wire.NewSet(NewAuthRouter)

type AuthRouter struct {
	Auth           *controllers.AuthController
	DemoMiddleware *middlewares.DemoMiddleware
}

func NewAuthRouter(Auth *controllers.AuthController, DemoMiddleware *middlewares.DemoMiddleware) *AuthRouter {
	return &AuthRouter{Auth: Auth, DemoMiddleware: DemoMiddleware}
}

func (a *AuthRouter) InitRoutes(r *gin.Engine) *gin.Engine {
	auth := r.Group("/api/auth").Use(a.DemoMiddleware.Ac())
	{
		auth.GET("/test", a.Auth.Login)
		auth.GET("/testv2/:id", a.Auth.Loginv2)
	}
	return r
}
