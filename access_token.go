package dingtalk

import (
	"fmt"
	"sync"
)

const (
	FMT_URL_GET_ACCESS_TOKEN = `https://oapi.dingtalk.com/gettoken?corpid=%s&corpsecret=%s`
)

var (
	accees_mutex sync.Mutex
)

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	BaseResponse
}

func (client *DTalkClient) GetAccessToken() (string, error) {
	accees_mutex.Lock()
	defer accees_mutex.Unlock()

	if client.accessToken == "" {
		at_err := client.RefreshAccessToken()
		if at_err != nil {
			return "", at_err
		}
	}
	return client.accessToken, nil
}

func (client *DTalkClient) RefreshAccessToken() error {
	req_url := fmt.Sprintf(FMT_URL_GET_ACCESS_TOKEN, client.CorpID, client.CorpSecret)

	//http Get
	at_resp := &AccessTokenResp{}
	get_err := HttpGetJson(req_url, at_resp)
	if get_err != nil {
		return get_err
	}

	//赋值accesstoken
	get_err = at_resp.CheckError()
	if get_err != nil {
		return get_err
	} else {
		client.accessToken = at_resp.AccessToken
		return nil
	}
}
