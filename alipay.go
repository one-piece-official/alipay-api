package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/one-piece-official/alipay-api/dto"
	"github.com/one-piece-official/alipay-api/pkg/method"
	"github.com/one-piece-official/alipay-api/utils"
)

const (
	version    = "1.0"
	charset    = "utf-8"
	signType   = "RSA2"
	format     = "json"
	gatewayURL = "https://openapi.alipay.com/gateway.do"
)

type Client struct {
	appID    string
	appSecret string
}

// NewClient 初始化支付宝客户端.
func NewClient(key, secret string) *Client {
	return &Client{
		appID:    key,
		appSecret: secret,
	}
}

func (c *Client) Query(request method.QueryRequest, response interface{}) (err error) {
	bizContentStr, _ := json.Marshal(request)
	params := dto.RequestBody{
		AppID:      c.appID,
		Format:     format,
		Charset:    charset,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    version,
		SignType:   signType,
		Method:     request.GetMethod(),
		BizContent: string(bizContentStr),
	}

	signString, err := c.composeParameterString(params)
	if err != nil {
		return err
	}

	sign, err := utils.RsaSignWithSha256Hex(signString, c.appSecret)
	if err != nil {
		return fmt.Errorf("rsa sign failed: %w", err)
	}
	params.Sign = sign

	httpClient := http.DefaultClient
	resp, err := utils.MakeHTTPClientGet(httpClient, gatewayURL, bytes.NewBuffer(nil))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	resByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(bytes.NewBuffer(resByte))
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return
	}

	return nil
}

func (c *Client) composeParameterString(params dto.RequestBody) (signString string, err error) {
	// 将结构体转换为 map
	bytesData, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("compose parameter string json marshal failed: %w", err)
	}

	requestDataMap := make(map[string]string)
	_ = json.Unmarshal(bytesData, &requestDataMap)

	// 遍历 map 将 key 取出来并按照 ascii 排序
	keys := make([]string, 10)
	for key := range requestDataMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %s", key, requestDataMap[key])
		signString += key + "=" + requestDataMap[key] + "&"
	}
	// 去掉最后一个 "&"
	signString = signString[:len(signString)-1]

	return signString, nil
}

func buildBizContent(params dto.RequestBody) string {
	var mc = make(map[string]string)
	byteBiz, _ := json.Marshal(mc)
	biz := string(byteBiz)
	return biz
}