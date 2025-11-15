package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterProviderSet = wire.NewSet(NewRouter, AuthRouterProviderSet)

func NewRouter(AuthRouter *AuthRouter) *gin.Engine {
	r := gin.Default()
	r = AuthRouter.InitRoutes(r)
	return r

}
