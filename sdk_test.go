package go_wechat_miniapp_sdk

import (
	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/config"
	"github.com/dgb8901/go-wechat-miniapp-sdk/service"
	"log"
	"testing"
)

func Test_sdk(t *testing.T) {
	cfg := &config.Cfg{
		AppId:         "",
		Secret:        "",
		Token:         "",
		AesKey:        "",
		MsgDataFormat: "JSON",
	}
	//redisConfig := config.NewInRedis(cfg,"10.19.9.16:6379","")
	memory := config.NewInMemory(cfg)
	wxaService := service.NewService(memory)
	userService := wxaService.GetUserService()

	session, err := userService.Jscode2Session("091Hdq100gkbBL10Dg300Xa1BF4Hdq16")

	log.Printf(common.ToJson(err))
	log.Printf(common.ToJson(session))

}
