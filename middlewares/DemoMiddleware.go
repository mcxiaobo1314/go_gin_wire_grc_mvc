package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var DemoMiddlewareProviderSet = wire.NewSet(NewDemoMiddleware)

type DemoMiddleware struct {
	Db *gorm.DB
}

func NewDemoMiddleware(Db *gorm.DB) *DemoMiddleware {
	return &DemoMiddleware{Db: Db}
}

func (d *DemoMiddleware) Ac() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("DemoMiddleware")
		//ctx.Abort() //终止请求
		ctx.Next() //继续请求

	}
}
