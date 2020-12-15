# alipay-api
支付宝 api 客户端

## 使用

```go
package main

import (
    "fmt"
    alipay_api "github.com/one-piece-official/alipay-api"
)

func main() {
	client := alipay_api.NewClient(
		"configs.GetConfig().Alipay.AppKey",
		"configs.GetConfig().Alipay.AppSecret",
		"configs.GetConfig().Alipay.URL",
	)
	res, err := client.Execute("alipay.user.account.device.info.query", map[string]interface{}{
		"device_type": "IMEI", "request_from": "request_from", "encrypt_type": "",
		"device_ids": []string{"IMEI"},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```

