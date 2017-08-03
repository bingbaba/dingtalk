package dingtalk

import (
	"encoding/json"
	"fmt"
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

func HttpJson(req *http.Request, v interface{}) error {
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", data)
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
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
