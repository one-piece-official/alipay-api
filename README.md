# alipay-api
支付宝 api 客户端

## 使用

```go
package main

import (
    "fmt"
    alipay "github.com/one-piece-official/alipay-api"
)

func main() {
	client := alipay.NewClient(
		"configs.GetConfig().Alipay.AppKey",
		"configs.GetConfig().Alipay.AppSecret",
		"configs.GetConfig().Alipay.URL",
	)
	
	var res method.UserAccountDeviceInfoResponse
	err := client.Query(&method.UserAccountDeviceInfoRequest{
		DeviceType:  "IMEI",
		RequestFrom: "requestfrom",
		EncryptType: "",
		DeviceIds:   []string{"111"},
	}, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.AlipayUserAccountDeviceInfoQueryResponse.DeviceInfos[0].DeviceLabel)
}
```

