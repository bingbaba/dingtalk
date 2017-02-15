package dingtalk

import (
	"sync"
	//"time"
	"strings"
	"time"
)

type MonitorMsg interface {
	Exception() bool
	GetID() string
	GetName() string
	GetRichNum() string
	GetRichUnit() string
	GetForm() [][2]string
	GetContent() string
	GetStartTime() int64
	GetAuthor() string
	GetUrl() string
}

type DefaultMonitorMsg struct {
}

func (m *DefaultMonitorMsg) Exception() bool {
	return true
}
func (m *DefaultMonitorMsg) GetID() string {
	return ""
}
func (m *DefaultMonitorMsg) GetName() string {
	return ""
}
func (m *DefaultMonitorMsg) GetRichNum() string {
	return ""
}
func (m *DefaultMonitorMsg) GetRichUnit() string {
	return ""
}
func (m *DefaultMonitorMsg) GetForm() [][2]string {
	return make([][2]string, 0)
}
func (m *DefaultMonitorMsg) GetContent() string {
	return ""
}
func (m *DefaultMonitorMsg) GetStartTime() int64 {
	return time.Now().Unix()
}
func (m *DefaultMonitorMsg) GetAuthor() string {
	return ""
}
func (m *DefaultMonitorMsg) GetUrl() string {
	return ""
}

type DingTalkConf struct {
	Group      string `json:"group" toml:"group" yaml:"group"`
	AgentID    string `json:"agentid" toml:"agentid" yaml:"agentid"`
	DepartID   string `json:"departion" toml:"departion" yaml:"departion"`
	CorpID     string `json:"corpid" toml:"corpid" yaml:"corpid"`
	CorpSecret string `json:"corpsecret" toml:"corpsecret" yaml:"corpsecret"`
}

type DingtalkTool struct {
	conf   *DingTalkConf
	client *DTalkClient
	m      sync.Mutex
}

func NewDingtalkTool(conf *DingTalkConf) *DingtalkTool {
	return &DingtalkTool{conf: conf}
}

func (dt *DingtalkTool) SendOA(msg MonitorMsg) error {
	oamd := &MsgDetail_content_oa{}
	oamd.MsgUrl = msg.GetUrl()
	oamd.Head.Text = msg.GetName()
	if msg.Exception() {
		oamd.Body.Title = msg.GetName() + "异常"
		oamd.Body.Content = msg.GetContent()
	} else {
		oamd.Body.Title = msg.GetName() + "正常"
		oamd.Head.BgColor = "FF258DEB"
		oamd.Body.Content = msg.GetContent()
	}

	form := append(msg.GetForm(), [2]string{
		"开始时间：", time.Unix(msg.GetStartTime(), 0).Format("01/02 15:04"),
	})
	oamd.Body.Form = ParseForm(form)

	oamd.Body.Rich.Num = msg.GetRichNum()
	oamd.Body.Rich.Unit = msg.GetRichUnit()
	oam := &MsgDetail_oa{
		Msgtype: "oa",
		OA:      oamd,
	}
	oamd.Body.Author = msg.GetAuthor()

	client, err := dt.getClient()
	if err != nil {
		return err
	}
	dtm := NewPartyOAMessage(dt.conf.AgentID, dt.conf.DepartID, oam)
	return client.SendToCompany(dtm)
}

func (dt *DingtalkTool) getClient() (*DTalkClient, error) {
	dt.m.Lock()
	defer dt.m.Unlock()

	if dt.client == nil {
		var cerr error
		dt.client, cerr = NewDTalkClient(dt.conf.CorpID, dt.conf.CorpSecret)
		if cerr != nil {
			return nil, cerr
		}
	}

	return dt.client, nil
}

func ParseForm(fa [][2]string) []KeyValue {
	form := make([]KeyValue, 0)

	for _, kv := range fa {
		if !strings.HasSuffix(kv[0], ":") && !strings.HasSuffix(kv[0], "：") {
			kv[0] = kv[0] + "："
		}
		form = append(form, KeyValue{
			kv[0], kv[1],
		})
	}

	return form
}

func MapToForm(m map[string]string) (ret [][2]string) {
	ret = make([][2]string, 0)
	for key, value := range m {
		ret = append(ret, [2]string{key, value})
	}
	return
}
