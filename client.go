package alipay_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/pkg/errors"
)

const (
	version  = "1.0"
	charset  = "utf-8"
	signType = "RSA2"
)

type Client interface {
	Execute(method string, query map[string]interface{}) (map[string]interface{}, error)
}

type client struct {
	appKey    string
	appSecret string
	url       string
}

func NewClient(appKey string, appSecret string, url string) Client {
	return &client{
		appKey:    appKey,
		appSecret: appSecret,
		url:       url,
	}
}

// nolint: noctx
func (c *client) Execute(method string, query map[string]interface{}) (mapData map[string]interface{}, err error) {
	params := c.commonParams()
	params["method"] = method
	bizContent := map[string]interface{}{}
	for key, value := range query {
		bizContent[key] = value
	}
	bizContentStr, _ := json.Marshal(bizContent)
	params["biz_content"] = string(bizContentStr)
	params["sign"] = sign(params, c.appSecret)
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, value)
	}
	resp, err := http.PostForm(c.url, urlParams)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.NewDecoder(bytes.NewBuffer(resByte)).Decode(&mapData)
	if err != nil {
		return map[string]interface{}{}, errors.Wrap(err, "alipay json decode fail")
	}

	return mapData, nil
}

func sign(params map[string]string, secret string) (sign string) {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	buffer := new(bytes.Buffer)
	for _, k := range keys {
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, params[k]))
	}
	s, i := buffer.String(), buffer.Len()
	sign, err := util.RsaSignWithSha256Hex(s[:i-1], secret)
	if err != nil {
		return
	}

	return sign
}

func (c *client) commonParams() map[string]string {
	return map[string]string{
		"app_id":    c.appKey,
		"charset":   charset,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"version":   version,
		"sign_type": signType,
	}
}
