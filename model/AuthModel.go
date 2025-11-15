package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var AuthModelProviderSet = wire.NewSet(NewAuthModel)

type AuthModel struct {
	Db *gorm.DB
}

func NewAuthModel(Db *gorm.DB) *AuthModel {
	return &AuthModel{
		Db: Db,
	}
}
