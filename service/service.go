package service

import (
	"sync"
)

const (
	// 获取access_token请求地址
	getAccessTokenUrl string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	// 重试次数
	Retries int = 3
	// 请求方法
	GET    string = "GET"
	POST   string = "POST"
	DELETE string = "DELETE"
	PUT    string = "PUT"
)

var lock sync.Mutex

type Service interface {
	// Get get请求
	Get(url string, params *map[string]interface{}, resp interface{}) error

	// Post post请求
	Post(url string, data *map[string]interface{}, resp interface{}) error

	// GetFile get请求,请求文件
	GetFile(url string, params *map[string]interface{}) ([]byte, error)

	// PostFile postFile请求
	PostFile(url string, data *map[string]interface{}) ([]byte, error)

	// GetAccessToken 获取access_token
	GetAccessToken() (string, error)

	// CheckSignature 验证消息的确来自微信服务器.
	CheckSignature(timestamp string, nonce string, signature string) bool
}
