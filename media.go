package dingtalk

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	MEDIATYPE_IMAGE      = "image"
	MEDIATYPE_VIOCE      = "voice"
	MEDIATYPE_FILE       = "file"
	FMT_URL_MEDIA_UPLOAD = "https://oapi.dingtalk.com/media/upload?access_token=%s&type=%s"
)

type MediaUploadResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	MediaType  string `json:"type"`
	MediaID    string `json:"media_id"`
	CreateTime int64  `json:"created_at"`
}

func (client *DTalkClient) UploadMedia(filename, mediatype string) (string, error) {
	accees_token, at_err := client.GetAccessToken()
	if at_err != nil {
		return "", at_err
	}

	urlStr := fmt.Sprintf(FMT_URL_MEDIA_UPLOAD, accees_token, mediatype)

	var param = map[string]string{
	// "filename": filepath.Base(filename),
	}
	req, err := newfileUploadRequest(urlStr, param, "media", filename)
	// req, err := http.NewRequest("POST", urlStr, bodyBuf)
	if err != nil {
		return "", err
	}

	resp := &MediaUploadResponse{}
	err = HttpJson(req, resp)
	if err != nil {
		return "", err
	}

	if resp.ErrMsg != "ok" {
		return "", errors.New(resp.ErrMsg)
	}

	return resp.MediaID, nil
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
