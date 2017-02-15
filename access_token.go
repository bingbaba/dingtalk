package dingtalk

import (
	"fmt"
)

const (
	FMT_URL_GET_ACCESS_TOKEN = `https://oapi.dingtalk.com/gettoken?corpid=%s&corpsecret=%s`
)

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	BaseResponse
}

func (client *DTalkClient) GetAccessToken() (string, error) {
	client.tokenM.Lock()
	defer client.tokenM.Unlock()

	if client.accessToken == "" {
		at_err := client.refreshAccessToken()
		if at_err != nil {
			return "", at_err
		}
	}
	return client.accessToken, nil
}

func (client *DTalkClient) RefreshAccessToken() error {
	client.tokenM.Lock()
	defer client.tokenM.Unlock()
	return client.refreshAccessToken()
}

func (client *DTalkClient) refreshAccessToken() error {
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
