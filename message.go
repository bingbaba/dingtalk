package dingtalk

type Message_text struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type PartyTextMessage struct {
	AgentID string `json:"agentid"`
	Toparty string `json:"toparty"`
	*Message_text
}

func NewMessage_text(content string) *Message_text {
	text := &Message_text{Msgtype: "text"}
	text.Text.Content = content
	return text
}

func NewPartyMessage(agentid, partyid, content string) *PartyTextMessage {
	return &PartyTextMessage{
		agentid,
		partyid,
		NewMessage_text(content),
	}
}
