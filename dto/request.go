package dto

type RequestBody struct {
	// 公共参数
	AppID     string `json:"app_id"`         // 应用 id
	Charset   string `json:"charset"`        // 请求使用的编码
	Format    string `json:"format"`         // 格式，选填
	Method    string `json:"method"`         // 接口名称
	SignType  string `json:"sign_type"`      // 生成签名的算法
	Sign      string `json:"sign,omitempty"` // 签名
	Timestamp string `json:"timestamp"`      // 时间
	Version   string `json:"version"`        // 版本

	// 请求参数集合
	BizContent string `json:"biz_content,omitempty"`
}

type UserAccountDeviceInfoRequest struct {
	DeviceType  string   `json:"device_type"`
	RequestFrom string   `json:"request_from"`
	EncryptType string   `json:"encrypt_type"`
	DeviceIDs   []string `json:"device_ids"`
}

func (req UserAccountDeviceInfoRequest) GetMethod() string {
	return "alipay.user.account.device.info.query"
}