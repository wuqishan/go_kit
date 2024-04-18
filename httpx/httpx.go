package httpx

import (
	"bytes"
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"net/url"
)

// Post post请求
// u url
// data map[string]interface{}
// header map[string]string
func Post(u string, data interface{}, header map[string]string) (res []byte, err error) {

	val, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(val))
	if err != nil {
		return nil, err
	}
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	if len(header) == 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get post请求
// u url
func Get(u string) (res []byte, err error) {

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PostForm ...
// u url
// data 参数
func PostForm(u string, data map[string]interface{}) (string, error) {
	tmp := url.Values{}
	for k, v := range data {
		tmp.Add(k, gconv.String(v))
	}
	resp, err := http.PostForm(u, tmp)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SimpleGet(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
