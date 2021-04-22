# go-wechat-miniapp-sdk

#### 关于

> 基于微信小程序相关接口封装，使用golang语言封装的一套微信小程序官方接口SDK
>
> 若在使用过程中出现bug,请到[issues](https://github.com/dgb8901/go-wechat-miniapp-sdk/issues)提问.

#### 版本

> 当前版本: [v1.0.0-alpha](https://github.com/dgb8901/go-wechat-miniapp-sdk/releases/tag/v1.0.0-alpha)

#### 运行环境

> `>= 1.14`

#### 安装方法

> 执行命令
>
> `go get github.com/dgb8901/go-wechat-miniapp-sdk`

#### 快速使用

##### 初始化SDK

```golang
package helper

import (
	"github.com/dgb8901/go-wechat-miniapp-sdk/config"
	"github.com/dgb8901/go-wechat-miniapp-sdk/service"
)

type WxaHelper struct {
	wxaService *service.WxaService
}

var wxaHelper = &WxaHelper{}

func Init() {
	wxaConfig := &config.WxaInMemoryConfig{
		AppId:         "AppId",
		Secret:        "Secret",
		Token:         "Token",
		AesKey:        "AesKey",
		MsgDataFormat: "DataFormat",
	}

	wxaService := &service.New(wxaConfig)
	wxaHelper.wxaService = wxaService
}

func GetWxaService() *service.WxaService {
	return wxaHelper.wxaService
}

```

##### 使用

```golang
// 获取用户service
userService := helper.GetWxaService().GetUserService()
// 根据用户service获取用户session
session, err := userService.Jscode2Session(jsCode)
```

#### 功能列表

* [登录|用户信息](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/service/user_service.go)
* [订阅消息](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/service/subscribe_msg_service.go)
* [客服消息](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/service/kf_service.go)
* [统一服务消息](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/service/uniform_message_service.go)
* [获取小程序码](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/service/qr_code_service.go)
* ...

#### 联系我

> 邮箱: dgb8901@163.com
>
> QQ: 770713275

#### 作者

> [dgb8901](https://github.com/dgb8901)

#### License

> MIT License, see [license file](https://github.com/dgb8901/go-wechat-miniapp-sdk/blob/main/License)
