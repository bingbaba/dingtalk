package dingtalk

import (
	"sync"
	"time"
)

const (
	REFRESH_TOKEN_INTERVAL = 7000
)

type DTalkClient struct {
	CorpID       string
	CorpSecret   string
	accessToken  string
	getTokenTime int64
	tokenM       sync.Mutex
}

func NewDTalkClient(corpid, corpSecret string) (*DTalkClient, error) {
	client := &DTalkClient{
		CorpID:     corpid,
		CorpSecret: corpSecret,
	}

	err := client.RefreshAccessToken()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			time.Sleep(REFRESH_TOKEN_INTERVAL * time.Second)
			if client.accessToken != "" {
				client.accessToken = ""
			}
		}
	}()

	return client, nil
}
