package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	METHOD_CORPCONVERSATION_ASYNC = "dingtalk.corp.message.corpconversation.asyncsend"
	URL_CORPCONVERSATION_ASYNC    = "https://eco.taobao.com/router/rest"
)

/*
type RequestCommonPara struct {
	Method    string `json:"method"`
	Session   string `json:"session"`
	TimeStamp string `json:"timestamp"`
	Format    string `json:"format,omitempty"`
	Version   string `jsoN:"v"`
	PartnerID string `json:"partner_id,omitempty"`
	Simplify  bool   `json:"simplify,omitempty"`
}

type RequestMessage struct {
	MsgType    string   `json:"msgtype"`
	AgentID    string   `json:"agent_id"`
	UserIDList []string `json:"userid_list"`
	DeptIDList []string `json:"dept_id_list"`
	ToAllUser  bool     `json:"to_all_user"`
}
*/

type MsgTarget struct {
	UserIDList []string
	DeptIDList []string
	ToAllUsers bool
}

type OAMsgContent struct {
	MsgUrl string            `json:"message_url"`
	Head   OAMsgContentHead  `json:"head"`
	Body   *OAMsgContentBody `json:"body"`
}

type OAMsgContentHead struct {
	BgColor string `json:"bgcolor"`
	Text    string `json:"text"`
}

type OAMsgContentBody struct {
	Title     string               `json:"title,omitempty"`
	Form      []KeyValue           `json:"form,omitempty"`
	Rich      OAMsgContentBodyRich `json:"rich,omitempty"`
	Content   string               `json:"content,omitempty"`
	Image     string               `json:"image,omitempty"`
	FileCount string               `json:"file_count,omitempty"`
	Author    string               `json:"author,omitempty"`
}
type OAMsgContentBodyRich struct {
	Num  string `json:"num"`
	Unit string `json:"unit"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OAMsgResponse struct {
	MethodResponse struct {
		Result *MsgResponse `json:"result"`
	} `json:"dingtalk_corp_message_corpconversation_asyncsend_response"`
	ErrorResp *ErrorResp `json:"error_response"`
}

type ErrorResp struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	ReqID string `json:"request_id"`
}

type MsgResponse struct {
	DingOpenErrCode int    `json:"ding_open_errcode"`
	ErrorMsg        string `json:"error_msg"`
	Success         bool   `json:"success"`
	TaskID          int    `json:"task_id"`
}

func (client *DTalkClient) SendOAMsg(agentid string, target *MsgTarget, content *OAMsgContent) error {
	accees_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return at_err
	}

	content_bytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	// content_str := string(content_bytes)
	// content_str = `{"message_url": "http://dingtalk.com","head": {"bgcolor": "FFBBBBBB","text": "头部标题"},"body": {"title": "正文标题","form": [{"key": "姓名:","value": "张三"},{"key": "爱好:","value": "打球、听音乐"}],"rich": {"num": "15.6","unit": "元"},"content": "大段文本大段文本大段文本大段文本大段文本大段文本大段文本大段文本大段文本大段文本大段文本大段文本","image": "@lADOADmaWMzazQKA","file_count": "3","author": "李四 "}}`

	param := url.Values{
		"method":      {METHOD_CORPCONVERSATION_ASYNC},
		"session":     {accees_token},
		"timestamp":   {time.Now().Format("2006-01-02 15:04:05")},
		"format":      {"json"},
		"v":           {"2.0"},
		"msgtype":     {"oa"},
		"agent_id":    {agentid},
		"userid_list": {strings.Join(target.UserIDList, ",")},
		// "dept_id_list": {strings.Join(target.DeptIDList, ",")},
		"to_all_user": {fmt.Sprintf("%v", target.ToAllUsers)},
		"msgcontent":  {string(content_bytes)},
	}

	req, err := http.NewRequest("POST", URL_CORPCONVERSATION_ASYNC, strings.NewReader(param.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp := &OAMsgResponse{}
	err = HttpJson(req, resp)
	if err != nil {
		return err
	}

	if resp.ErrorResp != nil {
		return errors.New(resp.ErrorResp.Msg)
	} else if !resp.MethodResponse.Result.Success {
		return errors.New(resp.MethodResponse.Result.ErrorMsg)
	}

	return nil
}
