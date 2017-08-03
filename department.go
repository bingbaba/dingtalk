package dingtalk

import (
	"errors"
	"fmt"
)

type Department struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Parentid        int    `json:"parentid"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	AutoAddUser     bool   `json:"autoAddUser"`
}

type DepartmentResp struct {
	BaseResponse
	Departments Departments `json:"department"`
}

type Departments []*Department

const (
	FMT_URL_GET_DEPARTLIST = `https://oapi.dingtalk.com/department/list?access_token=%s`
)

func (client *DTalkClient) GetAllDeparts() (Departments, error) {
	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return nil, at_err
	}

	resp := &DepartmentResp{}
	req_url := fmt.Sprintf(FMT_URL_GET_DEPARTLIST, access_token)
	get_err := HttpGetJson(req_url, resp)
	if get_err != nil {
		return nil, get_err
	}

	get_err = resp.CheckError()
	if get_err != nil {
		return nil, errors.New(resp.ErrMsg)
	} else {
		return resp.Departments, nil
	}
}

func (departments *Departments) GetDepartByName(depart_name string) *Department {
	for _, department := range []*Department(*departments) {
		if department.Name == depart_name {
			return department
		}
	}

	return nil
}
