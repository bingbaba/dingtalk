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

func (client *DTalkClient) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	depts, err := client.GetAllDeparts()
	if err != nil {
		return users, err
	}

	for _, dept := range depts {
		tmp_users, err := client.GetDeptUsers(dept.ID)
		if err != nil {
			return users, err
		}
		users = append(users, tmp_users...)
	}

	return users, nil
}

func (client *DTalkClient) GetDeptUsers(deptid int) ([]*User, error) {
	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return nil, at_err
	}

	resp := &UserResp{}

	req_url := fmt.Sprintf(FMT_URL_GET_USERLIST, access_token, deptid)
	get_err := HttpGetJson(req_url, resp)
	if get_err != nil {
		return nil, get_err
	}

	get_err = resp.CheckError()
	if get_err != nil {
		return nil, errors.New(resp.ErrMsg)
	} else {
		return resp.UserList, nil
	}
}
