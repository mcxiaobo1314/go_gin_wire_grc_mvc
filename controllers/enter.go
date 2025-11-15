package controllers

import "github.com/google/wire"

var ControllersProviderSet = wire.NewSet(AuthControllerProviderSet)
