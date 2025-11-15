package model

import "github.com/google/wire"

var EnterModelProviderSet = wire.NewSet(AuthModelProviderSet)
