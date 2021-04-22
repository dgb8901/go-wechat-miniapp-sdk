package service

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"go-wechat-miniapp-sdk/config"
	"go-wechat-miniapp-sdk/models/response"
	"log"
	"strings"
)

var services = make(map[string]interface{})

type wxaService struct {
	wxaConfig *config.WxaInMemoryConfig
}

func New(wxaConfig *config.WxaInMemoryConfig) *wxaService {
	return &wxaService{wxaConfig: wxaConfig}
}

func (wxa *wxaService) Get(url string, params *map[string]interface{}, resp interface{}) error {
	return wxa.execute(GET, url, params, resp)
}

func (wxa *wxaService) Post(url string, data *map[string]interface{}, resp interface{}) error {
	return wxa.execute(POST, url, data, resp)
}

func (wxa *wxaService) GetFile(url string, params *map[string]interface{}) ([]byte, error) {
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

func (wxa *wxaService) PostFile(url string, data *map[string]interface{}) ([]byte, error) {
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

func (wxa *wxaService) execute(method string, uri string, params *map[string]interface{}, resp interface{}) error {

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

func (wxa *wxaService) GetAccessToken() (string, error) {

	if !wxa.wxaConfig.IsAccessTokenExpired() {
		return wxa.wxaConfig.GetAccessToken(), nil
	}

	lock.Lock()
	defer lock.Unlock()

	url := fmt.Sprintf(getAccessTokenUrl, wxa.wxaConfig.GetAppId(), wxa.wxaConfig.GetSecret())

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

	wxa.wxaConfig.UpdateAccessToken(accessToken.AccessToken, accessToken.ExpiresIn)

	return accessToken.AccessToken, err
}

func (wxa *wxaService) GetWxaConfig() *config.WxaInMemoryConfig {
	return wxa.wxaConfig
}

func (wxa *wxaService) CheckSignature(timestamp string, nonce string, signature string) bool {
	return false
}

func (wxa *wxaService) SetWxaConfig(config *config.WxaInMemoryConfig) {
	wxa.wxaConfig = config
}

func (wxa *wxaService) GetUserService() *WxaUserService {
	userService := services["userService"]
	if userService == nil {
		userService = &WxaUserService{wxaService: wxa}
		services["userService"] = userService
	}
	service := userService.(*WxaUserService)
	return service
}

func (wxa *wxaService) GetSubscribeMsgService() *WxaSubscribeMsgService {
	subscribeMsgService := services["subscribeMsgService"]
	if subscribeMsgService == nil {
		subscribeMsgService = &WxaSubscribeMsgService{wxaService: wxa}
		services["subscribeMsgService"] = subscribeMsgService
	}
	service := subscribeMsgService.(*WxaSubscribeMsgService)
	return service
}

func (wxa *wxaService) GetKfService() *WxaKfService {
	kfService := services["kfService"]
	if kfService == nil {
		kfService = &WxaKfService{wxaService: wxa}
		services["kfService"] = kfService
	}
	service := kfService.(*WxaKfService)
	return service
}

func (wxa *wxaService) GetUniformMessageService() *UniformMessageService {
	uniformMessageService := services["uniformMessageService"]
	if uniformMessageService == nil {
		uniformMessageService = &UniformMessageService{wxaService: wxa}
		services["uniformMessageService"] = uniformMessageService
	}
	service := uniformMessageService.(*UniformMessageService)
	return service
}

func (wxa *wxaService) GetQrCodeService() *QrCodeService {
	qrCodeService := services["qrCodeService"]
	if qrCodeService == nil {
		qrCodeService = &QrCodeService{wxaService: wxa}
		services["qrCodeService"] = qrCodeService
	}
	service := qrCodeService.(*QrCodeService)
	return service
}
