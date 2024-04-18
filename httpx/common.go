package httpx

import (
	"bytes"
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	TimeoutDefault = time.Second * 5
)

// Conf ...
type Conf struct {
	Url     string
	Header  map[string]string
	Params  map[string]interface{}
	Timeout time.Duration
	Method  string
}

// Do ...
func Do(c *Conf) (res []byte, err error) {

	params, err := c.getParams()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(strings.ToUpper(c.Method), c.Url, bytes.NewBuffer(params))
	if err != nil {
		return nil, err
	}

	req = c.getHeaders(req)
	client := c.getClient()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Conf) getClient() *http.Client {
	if c.Timeout <= 0 {
		c.Timeout = TimeoutDefault // 默认超时时间
	}
	return &http.Client{Timeout: c.Timeout}
}

func (c *Conf) getHeaders(req *http.Request) *http.Request {
	if len(c.Header) > 0 {
		for k, v := range c.Header {
			req.Header.Set(k, v)
		}
	} else {
		req.Header.Set("Content-Type", "application/json") // 默认header
	}

	return req
}

func (c *Conf) getParams() ([]byte, error) {
	var (
		p   []byte
		err error
	)
	if len(c.Params) > 0 {
		switch c.Method {
		case http.MethodGet:
			var tmp *url.URL
			tmp, err = url.Parse(c.Url)
			if err != nil {
				return nil, err
			}
			val := tmp.Query()
			for k, v := range c.Params {
				val.Set(k, gconv.String(v))
			}
			tmp.RawQuery = val.Encode()
			c.Url = tmp.String()
		case http.MethodPost:
			p, err = json.Marshal(c.Params)
		default:
			// todo...
		}
	}
	return p, err
}
