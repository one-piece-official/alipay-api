package method

type UserAccountDeviceInfoRequest struct {
	DeviceType  string   `json:"device_type"`
	RequestFrom string   `json:"request_from"`
	EncryptType string   `json:"encrypt_type"`
	DeviceIds   []string `json:"device_ids"`
}

func (req UserAccountDeviceInfoRequest) GetMethod() string {
	return "alipay.user.account.device.info.query"
}

type UserAccountDeviceInfoResponse struct {
	AlipayUserAccountDeviceInfoQueryResponse AlipayUserAccountDeviceInfoQueryResponse `json:"alipay_user_account_device_info_query_response"`
	Sign                                     string                                    `json:"sign"`
}

type DeviceInfo struct {
	DeviceId string `json:"device_id"`
	DeviceLabel string `json:"device_label"`
}

type AlipayUserAccountDeviceInfoQueryResponse struct {
	Code        string        `json:"code"`
	Msg         string        `json:"msg"`
	ResultCode  string        `json:"result_code"`
	DeviceInfos []*DeviceInfo `json:"device_infos"`
}
