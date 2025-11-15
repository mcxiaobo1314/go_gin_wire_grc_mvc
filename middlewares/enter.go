package middlewares

import "github.com/google/wire"

var EnterMiddlewareProviderSet = wire.NewSet(DemoMiddlewareProviderSet)
