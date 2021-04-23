package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/dgb8901/go-wechat-miniapp-sdk/config"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/response"
	"log"
	"strings"
)

var services = make(map[string]interface{})

type WxaService struct {
	memory *config.WxaInMemoryConfig
	redis  *config.WxaInRedisConfig
	store  string
}

// 基于内存管理 access_token
func NewInMemory(wxaConfig *config.WxaInMemoryConfig) *WxaService {
	return &WxaService{memory: wxaConfig, store: "memory"}
}

// 基于Redis管理 access_token
func NewInRedis(wxaConfig *config.WxaInRedisConfig) *WxaService {
	return &WxaService{redis: wxaConfig, store: "redis"}
}

func (wxa *WxaService) Get(url string, params *map[string]interface{}, resp interface{}) error {
	return wxa.execute(GET, url, params, resp)
}

func (wxa *WxaService) Post(url string, data *map[string]interface{}, resp interface{}) error {
	return wxa.execute(POST, url, data, resp)
}

func (wxa *WxaService) GetFile(url string, params *map[string]interface{}) ([]byte, error) {
	var err error
	if strings.Contains(url, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token:" + url)
		log.Printf("The uri can't concat symbol access_token:%s", url)
		return nil, err
	}

	accessToken, err := wxa.GetAccessToken()
	if err != nil {
		return nil, err
	}
	if strings.Contains(url, "?") {
		url = url + "&access_token=" + accessToken
	} else {
		url = url + "?access_token=" + accessToken
	}

	request := httplib.NewBeegoRequest(url, GET)
	request.Retries(Retries)

	if params != nil {
		for k, v := range *params {
			request.Param(k, v.(string))
		}
	}
	resp, err := request.Bytes()

	return resp, err
}

func (wxa *WxaService) PostFile(url string, data *map[string]interface{}) ([]byte, error) {
	var err error
	if strings.Contains(url, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token:%s" + url)
		log.Printf("The uri can't concat symbol access_token:%s", url)
		return nil, err
	}

	accessToken, err := wxa.GetAccessToken()
	if err != nil {
		return nil, err
	}
	if strings.Contains(url, "?") {
		url = url + "&access_token=" + accessToken
	} else {
		url = url + "?access_token=" + accessToken
	}

	request := httplib.NewBeegoRequest(url, POST)
	request.Retries(Retries)

	if data != nil {
		request, _ = request.JSONBody(data)
	}

	resp, err := request.Bytes()
	return resp, err
}

func (wxa *WxaService) execute(method string, uri string, params *map[string]interface{}, resp interface{}) error {

	var err error
	if strings.Contains(uri, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token:%s" + uri)
		log.Printf("The uri can't concat symbol access_token:%s", uri)
		return err
	}

	accessToken, err := wxa.GetAccessToken()
	if err != nil {
		return err
	}
	if strings.Contains(uri, "?") {
		uri = uri + "&access_token=" + accessToken
	} else {
		uri = uri + "?access_token=" + accessToken
	}

	request := httplib.NewBeegoRequest(uri, method)
	request.Retries(Retries)

	if params != nil {
		if method == GET {
			for k, v := range *params {
				request.Param(k, v.(string))
			}
		} else if method == POST {
			request, _ = request.JSONBody(params)
		}
	}
	err = request.ToJSON(&resp)

	if err != nil {
		err = errors.New("request failure: %s" + err.Error())
		log.Printf("request failure: %s", err.Error())
		return err
	}
	return nil
}

// 更新access_token
func (wxa *WxaService) updateAccessToken(accessToken string, expiresInSeconds int64) {
	if wxa.store == "memory" {
		wxa.memory.UpdateAccessToken(accessToken, expiresInSeconds)
	} else if wxa.store == "redis" {
		wxa.redis.UpdateAccessToken(accessToken, expiresInSeconds)
	}
}

// 获取access_token
func (wxa *WxaService) GetAccessToken() (string, error) {

	if wxa.store == "memory" {
		if !wxa.memory.IsAccessTokenExpired() {
			return wxa.memory.GetAccessToken(), nil
		}
	} else if wxa.store == "redis" {
		if !wxa.redis.IsAccessTokenExpired() {
			return wxa.redis.GetAccessToken(), nil
		}
	}

	lock.Lock()
	defer lock.Unlock()

	cfg := wxa.GetConfig()
	url := fmt.Sprintf(getAccessTokenUrl, cfg.AppId, cfg.Secret)

	request := httplib.Get(url)
	request.Retries(3)
	accessToken := &response.WxaAccessToken{}
	err := request.ToJSON(accessToken)
	if err != nil {
		err = errors.New("request failure：%s" + err.Error())
		log.Printf("request failure：%s", err.Error())
		return "", err
	}
	if accessToken.ErrCode != 0 {
		err = errors.New(accessToken.ErrMsg)
		log.Printf(accessToken.ErrMsg)
		return "", err
	}

	wxa.updateAccessToken(accessToken.AccessToken, accessToken.ExpiresIn)

	return accessToken.AccessToken, err
}

func (wxa *WxaService) GetConfig() (cfg *config.Config) {

	if wxa.store == "memory" {
		return wxa.memory.GetConfig()
	} else if wxa.store == "redis" {
		return wxa.redis.GetConfig()
	}
	return nil
}

func (wxa *WxaService) CheckSignature(timestamp string, nonce string, signature string) bool {
	return false
}

func (wxa *WxaService) GetUserService() *wxaUserService {
	userService := services["userService"]
	if userService == nil {
		userService = &wxaUserService{wxaService: wxa}
		services["userService"] = userService
	}
	service := userService.(*wxaUserService)
	return service
}

func (wxa *WxaService) GetSubscribeMsgService() *wxaSubscribeMsgService {
	subscribeMsgService := services["subscribeMsgService"]
	if subscribeMsgService == nil {
		subscribeMsgService = &wxaSubscribeMsgService{wxaService: wxa}
		services["subscribeMsgService"] = subscribeMsgService
	}
	service := subscribeMsgService.(*wxaSubscribeMsgService)
	return service
}

func (wxa *WxaService) GetKfService() *wxaKfService {
	kfService := services["kfService"]
	if kfService == nil {
		kfService = &wxaKfService{wxaService: wxa}
		services["kfService"] = kfService
	}
	service := kfService.(*wxaKfService)
	return service
}

func (wxa *WxaService) GetUniformMessageService() *wxaUniformMessageService {
	uniformMessageService := services["uniformMessageService"]
	if uniformMessageService == nil {
		uniformMessageService = &wxaUniformMessageService{wxaService: wxa}
		services["uniformMessageService"] = uniformMessageService
	}
	service := uniformMessageService.(*wxaUniformMessageService)
	return service
}

func (wxa *WxaService) GetQrCodeService() *wxaQrCodeService {
	qrCodeService := services["qrCodeService"]
	if qrCodeService == nil {
		qrCodeService = &wxaQrCodeService{wxaService: wxa}
		services["qrCodeService"] = qrCodeService
	}
	service := qrCodeService.(*wxaQrCodeService)
	return service
}
