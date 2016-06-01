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
	BaseResponse
	UserList []*User `json:"userlist"`
}

func (client *DTalkClient) GetAlluser() ([]*User, error) {
	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return nil, at_err
	}

	resp := &UserResp{}

	//获取 department_id = 1的所有员工
	req_url := fmt.Sprintf(FMT_URL_GET_USERLIST, access_token, 1)
	get_err := HttpGetJson(req_url, resp)
	if get_err != nil {
		return nil, get_err
	}

	//检查返回错误
	get_err = resp.CheckError()
	if get_err != nil {
		return nil, errors.New(resp.ErrMsg)
	} else {
		return resp.UserList, nil
	}
}
