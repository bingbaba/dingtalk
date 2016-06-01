package dingtalk

import (
	"errors"
	"fmt"
)

type BaseResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (resp *BaseResponse) CheckError() error {
	//赋值accesstoken
	if resp.ErrCode != 0 || resp.ErrMsg != "ok" {
		return errors.New(fmt.Sprintf("errcode:%d,errmsg:%s", resp.ErrCode, resp.ErrMsg))
	} else {
		return nil
	}
}
