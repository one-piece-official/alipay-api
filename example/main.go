package main

import (
	"fmt"

	"github.com/one-piece-official/alipay-api"
	dto "github.com/one-piece-official/alipay-api/dto"
)

const APPID = "your app id"

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
your private key
-----END RSA PRIVATE KEY-----
`)

func main() {
	fmt.Println("Hello")

	client := alipay.NewClient(APPID, string(privateKey))

	var res dto.UserAccountDeviceInfoResponse
	err := client.Query(&dto.UserAccountDeviceInfoRequest{
		DeviceType:  "IMEI",
		RequestFrom: "requestfrom",
		EncryptType: "",
		DeviceIDs:   []string{"111"},
	}, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.AlipayUserAccountDeviceInfoQueryResponse.DeviceInfos[0].DeviceLabel)
}
