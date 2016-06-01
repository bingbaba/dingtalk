package dingtalk

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	httpclient *http.Client
)

func init() {
	timeout := 30 * time.Second
	httpclient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: 90 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			MaxIdleConnsPerHost:   100,
		},
	}
}

func HttpGetJson(req_url string, obj interface{}) error {
	//http GET
	resp, client_err := httpclient.Get(req_url)
	if client_err != nil {
		return client_err
	}
	defer resp.Body.Close()

	//result
	resp_body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		return read_err
	}

	//json Unmarshal
	json_err := json.Unmarshal(resp_body, obj)
	if json_err != nil {
		return json_err
	} else {
		return nil
	}
}

func HttpPostJson(req_url string, body_i, obj interface{}) error {
	body, json_err := json.Marshal(body_i)
	if json_err != nil {
		return json_err
	}

	//fmt.Printf("URL:%s\n", req_url)
	//fmt.Printf("BODY:%s\n", string(body))

	req, req_err := http.NewRequest("POST", req_url, bytes.NewReader(body))
	if req_err != nil {
		return req_err
	}
	req.Header.Set("Content-Type", "application/json")

	//http GET
	resp, client_err := httpclient.Do(req)
	if client_err != nil {
		return client_err
	}
	defer resp.Body.Close()

	//result
	resp_body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		return read_err
	}

	//fmt.Printf("RESP:%s\n", string(resp_body))

	//json Unmarshal
	json_err = json.Unmarshal(resp_body, obj)
	if json_err != nil {
		return json_err
	} else {
		return nil
	}
}
