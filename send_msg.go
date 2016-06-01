package dingtalk

import (
	"errors"
	"fmt"
)

const (
	FMT_URL_SEND_MESSAGE = `https://oapi.dingtalk.com/message/send?access_token=%s`
)

func (client *DTalkClient) SendToCompany(message interface{}) error {
	req_url, get_err := client.getSendCompanyUrl()
	if get_err != nil {
		return get_err
	}

	resp := &BaseResponse{}

	//获取 department_id = 1的所有员工
	get_err = HttpPostJson(req_url, message, resp)
	if get_err != nil {
		return get_err
	}

	//检查返回错误
	get_err = resp.CheckError()
	if get_err != nil {
		return errors.New(resp.ErrMsg)
	} else {
		return nil
	}
}

func (client *DTalkClient) getSendCompanyUrl() (string, error) {
	accees_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return "", at_err
	}
	return fmt.Sprintf(FMT_URL_SEND_MESSAGE, accees_token), nil
}
