package config

import "github.com/google/wire"

var EnterConfigProviderSet = wire.NewSet(MysqlProviderSet, RedisProviderSet, NewConfig)
