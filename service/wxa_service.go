package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/httplib"
	"github.com/dgb8901/go-wechat-miniapp-sdk/config"
	"github.com/dgb8901/go-wechat-miniapp-sdk/models/response"
)

type WxaService struct {
	cfg      config.CfgInterface
	services map[string]interface{}
}

func NewInMemoryService(cfg *config.Config) *WxaService {
	return &WxaService{
		cfg:      config.NewInMemory(cfg),
		services: make(map[string]interface{}),
	}
}

func NewInRedisService(cfg *config.Config) *WxaService {
	return &WxaService{
		cfg:      config.NewInRedis(cfg),
		services: make(map[string]interface{}),
	}
}

func (s *WxaService) Get(url string, params *map[string]interface{}, resp interface{}) error {
	return s.execute(GET, url, params, resp)
}

func (s *WxaService) Post(url string, data *map[string]interface{}, resp interface{}) error {
	return s.execute(POST, url, data, resp)
}

func (s *WxaService) GetFile(url string, params *map[string]interface{}) ([]byte, error) {
	var err error
	if strings.Contains(url, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token: " + url)
		return nil, err
	}

	accessToken, err := s.GetAccessToken()
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

func (s *WxaService) PostFile(url string, data *map[string]interface{}) ([]byte, error) {
	var err error
	if strings.Contains(url, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token: " + url)
		return nil, err
	}

	accessToken, err := s.GetAccessToken()
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

func (s *WxaService) execute(method string, uri string, params *map[string]interface{}, resp interface{}) error {

	var err error
	if strings.Contains(uri, "access_token=") {
		err = errors.New("The uri can't concat symbol access_token: " + uri)
		return err
	}

	accessToken, err := s.GetAccessToken()
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
		return err
	}
	return nil
}

// GetAccessToken 获取access_token
func (s *WxaService) GetAccessToken() (string, error) {

	if !s.cfg.IsAccessTokenExpired() {
		return s.cfg.GetAccessToken(), nil
	}

	lock.Lock()
	defer lock.Unlock()

	url := fmt.Sprintf(getAccessTokenUrl, s.cfg.GetAppId(), s.cfg.GetSecret())

	request := httplib.Get(url)
	request.Retries(3)
	accessToken := &response.WxaAccessToken{}
	err := request.ToJSON(accessToken)
	if err != nil {
		err = errors.New("request failure: %s" + err.Error())
		return "", err
	}
	if accessToken.ErrCode != 0 {
		err = errors.New(accessToken.ErrMsg)
		return "", err
	}

	s.cfg.UpdateAccessToken(accessToken.AccessToken, accessToken.ExpiresIn)

	return accessToken.AccessToken, err
}

func (s *WxaService) CheckSignature(timestamp string, nonce string, signature string) bool {
	// TODO check signature
	return false
}

func (s *WxaService) GetUserService() *WxaUserService {
	userService := s.services["userService"]
	if userService == nil {
		userService = newWxaUserService(s)
		s.services["userService"] = userService
	}
	service := userService.(*WxaUserService)
	return service
}

func (s *WxaService) GetSubscribeMsgService() *WxaSubscribeMsgService {
	subscribeMsgService := s.services["subscribeMsgService"]
	if subscribeMsgService == nil {
		subscribeMsgService = newWxaSubscribeMsgService(s)
		s.services["subscribeMsgService"] = subscribeMsgService
	}
	service := subscribeMsgService.(*WxaSubscribeMsgService)
	return service
}

func (s *WxaService) GetKfService() *WxaKfService {
	kfService := s.services["kfService"]
	if kfService == nil {
		kfService = newWxaKfService(s)
		s.services["kfService"] = kfService
	}
	service := kfService.(*WxaKfService)
	return service
}

func (s *WxaService) GetUniformMessageService() *WxaUniformMessageService {
	uniformMessageService := s.services["uniformMessageService"]
	if uniformMessageService == nil {
		uniformMessageService = newWxaUniformMessageService(s)
		s.services["uniformMessageService"] = uniformMessageService
	}
	service := uniformMessageService.(*WxaUniformMessageService)
	return service
}

func (s *WxaService) GetQrCodeService() *WxaQrCodeService {
	qrCodeService := s.services["qrCodeService"]
	if qrCodeService == nil {
		qrCodeService = newWxaQrCodeService(s)
		s.services["qrCodeService"] = qrCodeService
	}
	service := qrCodeService.(*WxaQrCodeService)
	return service
}
