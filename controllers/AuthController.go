package controllers

import (
	"gin_gorm/config"
	"gin_gorm/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/wire"
)

var AuthControllerProviderSet = wire.NewSet(NewAuthController)

type AuthController struct {
	AuthModel   *model.AuthModel
	RedisClient *redis.Client
	AppConfig   *config.Config
}

func NewAuthController(AuthModel *model.AuthModel, RedisClient *redis.Client, AppConfig *config.Config) *AuthController {
	return &AuthController{
		AuthModel:   AuthModel,
		RedisClient: RedisClient,
		AppConfig:   AppConfig,
	}
}

func (a *AuthController) Login(c *gin.Context) {
	type Params struct {
		Id string `form:"id" tag:"xxx" binding:"required"`
	}
	log.Println("AppConfig:", a.AppConfig)
	var params Params
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "xxxx:" + params.Id})
}

func (a *AuthController) Loginv2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is v2"})
}
