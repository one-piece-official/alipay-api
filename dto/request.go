package dto

type RequestBody struct {
	// 公共参数
	AppID     string `json:"app_id" url:"app_id" structs:"app_id"`               // 应用 id
	Charset   string `json:"charset" url:"charset" structs:"charset"`            // 请求使用的编码
	Format    string `json:"format" url:"format" structs:"format"`               // 格式，选填
	Method    string `json:"method" url:"method" structs:"method"`               // 接口名称
	SignType  string `json:"sign_type" url:"sign_type" structs:"sign_type"`      // 生成签名的算法
	Sign      string `json:"sign,omitempty" url:"sign" structs:"sign,omitempty"` // 签名
	Timestamp string `json:"timestamp" url:"timestamp" structs:"timestamp"`      // 时间
	Version   string `json:"version" url:"version" structs:"version"`            // 版本

	// 请求参数集合
	BizContent string `json:"biz_content,omitempty" url:"biz_content,omitempty" structs:"biz_content,omitempty"`
}

type UserAccountDeviceInfoRequest struct {
	DeviceType  string   `json:"device_type"`  // 设备类型，IMEI、IDFA、MOBILE(大小写敏感）
	RequestFrom string   `json:"request_from"` // 一般代表调用的合作机构名称，可写简称，大小写敏感
	EncryptType string   `json:"encrypt_type"` // 设备id的加密方式，如没有加密，可以不传。一般 MD5 即可满足需求
	DeviceIDs   []string `json:"device_ids"`   // IDFA 或者 IMEI 号数组。同一笔请求中，数组中只能是 IDFA 或者 IMEI,不能既有 IMEI，又有 IDFA
}

func (req *UserAccountDeviceInfoRequest) GetMethod() string {
	return "alipay.user.account.device.info.query"
}
