package dingtalk

import (
	"errors"
	"fmt"
)

const (
	FMT_URL_GET_USERLIST = `https://oapi.dingtalk.com/user/simplelist?access_token=%s&department_id=%d`
)

type User struct {
	Name   string `json:"name"`
	UserID string `json:"userid"`
}

type UserResp struct {
	ErrCode  int     `json:"errcode"`
	ErrMsg   string  `json:"errmsg"`
	UserList []*User `json:"userlist"`
}

func (client *DTalkClient) GetAlluser() ([]*User, error) {
	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return nil, at_err
	}

	user_resp := &UserResp{}

	//获取 department_id = 1的所有员工
	req_url := fmt.Sprintf(FMT_URL_GET_USERLIST, access_token, 1)
	get_err := HttpGetJson(req_url, user_resp)
	if get_err != nil {
		return nil, get_err
	}

	if user_resp.ErrCode != 0 {
		return nil, errors.New(user_resp.ErrMsg)
	} else {
		return user_resp.UserList, nil
	}
}
