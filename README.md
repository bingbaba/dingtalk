# dingtalk

## EXAMPLE
SendOA By DingTalkTool:
```golang
package main

import (
    "github.com/bingbaba/dingtalk"
)

func main() {
    conf := &dingtalk.DingTalkConf{
        AgentID:    "12632423",
        DepartID:   "83463432",
        CorpID:     "ding1238734234234f",
        CorpSecret: "HAsdfahsdfwe-DFHSF4SDH2FJSDFas3jdhf8aksjdhfkas3dGHSDFS",
    }

    dtt := dingtalk.NewDingtalkTool(conf)
    err := dtt.SendOA(&Message{})
    if err != nil {
        panic(err)
    }
}

type Message struct {
    *dingtalk.DefaultMonitorMsg
}

func (m *Message) GetForm() [][2]string {
    return [][2]string{
        [2]string{"FILED1", "VALUE1"},
        [2]string{"FILED2", "VALUE2"},
    }
}
func (m *Message) GetName() string {
    return "TEST"
}
func (m *Message) GetContent() string {
    return "content......"
}
func (m *Message) GetRichNum() string {
    return "1.5"
}
func (m *Message) GetRichUnit() string {
    return "RMB"
}
func (m *Message) GetAuthor() string {
    return "jack"
}

```