/**
** @创建时间: 2021/1/8 11:03 下午
** @作者　　: return
** @描述　　:
 */
package wechatEasySdk

import (
	"io/ioutil"
	"reflect"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 三方平台信息
 * @Date 2021/4/20 14:57:38
 * @Param
 * @return
 **/

type openBaseOptions struct {
	SpAppid               string `json:"sp_appid"` // 服务商应用ID
	SpMchid               string `json:"sp_mchid"` // 服务商户号
	AppId                 string `json:"appid"`
	AppSecret             string `json:"app_secret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
	Aeskey                string `json:"aeskey"`
	V3key                 string `json:"v3key"`
	RedirectUrl           string `json:"redirect_url"` // 开放平台验证票据回调url
	GatewayHost           string `json:"gateway_host"`
	AppCertPath           string `json:"app_cert_path"`
	PrivateKey            string `json:"private_key"`
	PublicKey             string `json:"public_key"`
	V3PublicKey           string `json:"v3_public_key"`
	WechatpaySerial       string `json:"wechatpay_serial"`
}

type baseOptions struct {
	AppId    string `json:"appid"`
	SubAppid string `json:"sub_appid"`
	MchId    string `json:"mch_id"`
	SubMchId string `json:"sub_mch_id"`
	Key      string `json:"key"`
}

var options *baseOptions

var openOptions *openBaseOptions

func NewOptions(params map[string]string) baseOptions {

	options = &baseOptions{
		AppId:    params["appid"],
		SubAppid: params["sub_appid"],
		MchId:    params["mch_id"],
		SubMchId: params["sub_mch_id"],
		Key:      params["key"],
	}

	return *options

}

func SetOption(key string, val string) {
	oPoint := reflect.ValueOf(options)
	field := oPoint.Elem().FieldByName(key)
	field.SetString(val)
}

func Options() *baseOptions {
	if options != nil {
		return options
	}
	return nil
}

func NewOpenOptions(params map[string]string) openBaseOptions {

	openOptions = &openBaseOptions{
		SpAppid:               params["spAppid"],
		SpMchid:               params["spMchid"],
		AppId:                 params["appId"],
		AppSecret:             params["appSecret"],
		Aeskey:                params["aesKey"],
		V3key:                 params["v3Key"],
		ComponentVerifyTicket: params["componentVerifyTicket"],
		RedirectUrl:           params["redirectUrl"],
		GatewayHost:           params["gatewayHost"],
		WechatpaySerial:       params["wechatpaySerial"],
	}

	if params["appCertPath"] != "" {

		openOptions.AppCertPath = params["appCertPath"]

		privateData, err := ioutil.ReadFile(params["appCertPath"] + "/wechat_private_key.pem")
		if err != nil {
			panic("读取私钥出错，文件不存在！")
		}

		openOptions.PrivateKey = string(privateData)

		publicData, err := ioutil.ReadFile(params["appCertPath"] + "/wechat_public_key.pem")
		if err != nil {
			panic("读取公钥钥出错，文件不存在！")
		}

		openOptions.PublicKey = string(publicData)

		v3publicData, err := ioutil.ReadFile(params["appCertPath"] + "/v3_public_key.pem")
		if err != nil {
			panic("读取公钥钥出错，文件不存在！")
		}

		openOptions.V3PublicKey = string(v3publicData)

	}

	return *openOptions

}

func SetOpenOption(key string, val string) {
	oPoint := reflect.ValueOf(openOptions)
	field := oPoint.Elem().FieldByName(key)
	field.SetString(val)
}

func OpenOptions() *openBaseOptions {
	if openOptions != nil {
		return openOptions
	}
	return nil
}
