package core

import "github.com/google/wire"

var CoreProviderSet = wire.NewSet(NewHServer, NewHGrpcServer, NewHGrpcClient)
