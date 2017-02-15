package dingtalk

import (
	"errors"
)

var (
	DefalutPool *DingTalkPool = &DingTalkPool{clients: make(map[string]*DingtalkTool)}
)

type DingTalkPool struct {
	clients map[string]*DingtalkTool
}

func NewDingTalkPool(dtcs []*DingTalkConf) (dtp *DingTalkPool, err error) {
	dtp = &DingTalkPool{
		clients: make(map[string]*DingtalkTool),
	}

	for _, conf := range dtcs {
		dttool, found := dtp.clients[conf.Group]
		if found {
			err = errors.New("the dingtalk id repeat: " + conf.Group)
			return
		}

		dttool = NewDingtalkTool(conf)
		dtp.clients[conf.Group] = dttool
	}

	return
}

func (dtp *DingTalkPool) SendOA(id string, msg MonitorMsg) (err error) {
	client, ok := dtp.clients[id]
	if !ok {
		err = errors.New("not found the " + id + " dingtalk client!")
		return
	}

	return client.SendOA(msg)
}

func InitDingtalk(dtcs []*DingTalkConf) (err error) {
	DefalutPool, err = NewDingTalkPool(dtcs)
	return
}

func SendOA(id string, msg MonitorMsg) error {
	return DefalutPool.SendOA(id, msg)
}
