package go_wechat_miniapp_sdk

import (
	"log"
	"testing"

	"github.com/dgb8901/go-wechat-miniapp-sdk/common"
	"github.com/dgb8901/go-wechat-miniapp-sdk/config"
	"github.com/dgb8901/go-wechat-miniapp-sdk/service"
)

func Test_sdk(t *testing.T) {
	cfg := &config.Config{
		AppId:         "",
		Secret:        "",
		Token:         "",
		AesKey:        "",
		MsgDataFormat: "JSON",
	}

	wxaService := service.NewInMemoryService(cfg)
	userService := wxaService.GetUserService()

	session, err := userService.Jscode2Session("091Hdq100gkbBL10Dg300Xa1BF4Hdq16")

	log.Printf(common.ToJson(err))
	log.Printf(common.ToJson(session))

}
