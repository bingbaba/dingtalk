package dingtalk

type Message_PartyText struct {
	AgentID string `json:"agentid"`
	Toparty string `json:"toparty"`
	*MsgDetail_text
}

type MsgDetail_text struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type Message_PartyOA struct {
	AgentID string `json:"agentid"`
	Toparty string `json:"toparty"`
	*MsgDetail_oa
}

type MsgDetail_oa struct {
	Msgtype string                `json:"msgtype"`
	OA      *MsgDetail_content_oa `json:"oa"`
}

type MsgDetail_content_oa struct {
	MsgUrl string `json:"message_url"`
	Head   struct {
		BgColor string `json:"bgcolor"`
		Text    string `json:"text"`
	} `json:"head"`
	Body struct {
		Title string     `json:"title,omitempty"`
		Form  []KeyValue `json:"form,omitempty"`
		Rich  struct {
			Num  string `json:"num"`
			Unit string `json:"unit"`
		} `json:"rich,omitempty"`
		Content   string `json:"content,omitempty"`
		Image     string `json:"image,omitempty"`
		FileCount string `json:"file_count,omitempty"`
		Author    string `json:"author,omitempty"`
	} `json:"body"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewMessage_text(content string) *MsgDetail_text {
	text := &MsgDetail_text{Msgtype: "text"}
	text.Text.Content = content
	return text
}

func NewPartyMessage(agentid, partyid, content string) *Message_PartyText {
	return &Message_PartyText{
		agentid,
		partyid,
		NewMessage_text(content),
	}
}

func NewPartyOAMessage(agentid, partyid string, oa_msg *MsgDetail_oa) *Message_PartyOA {
	return &Message_PartyOA{
		agentid,
		partyid,
		oa_msg,
	}
}
