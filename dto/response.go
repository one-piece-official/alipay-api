package dto

type ResponseBody struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
		ID     string `json:"id"`
	} `json:"data"`
}

type UserAccountDeviceInfoResponse struct {
	AlipayUserAccountDeviceInfoQueryResponse AlipayUserAccountDeviceInfoQueryResponse `json:"alipay_user_account_device_info_query_response"`
	Sign                                     string                                   `json:"sign"`
}

type AlipayUserAccountDeviceInfoQueryResponse struct {
	Code        string        `json:"code"`
	Msg         string        `json:"msg"`
	ResultCode  string        `json:"result_code"`
	DeviceInfos []*DeviceInfo `json:"device_infos"`
}

type DeviceInfo struct {
	DeviceID    string `json:"device_id"`
	DeviceLabel string `json:"device_label"`
}

