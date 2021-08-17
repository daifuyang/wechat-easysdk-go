/**
** @创建时间: 2021/1/4 11:19 下午
** @作者　　: return
** @描述　　:
 */
package pay

import (
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"testing"
)

func init() {
	params := map[string]string{
		"spAppid":     "wx842f10c69bb48e5b",
		"spMchid":     "1605269485",
		"appId":       "wxa95825ecf5e840e6",
		"appSecret":   "c733bbd5643a248fe19e7af3a837fcf1",
		"aesKey":      "codecloud2021codecloud2021codecloud20212021=",
		"notify_url":  "https://console.mashangdian.cn/api/v1/wechat/auth_redirect",
		"gatewayHost": "https://api.mch.weixin.qq.com",
		"appCertPath": "../pem",
	}

	wechatEasySdk.NewOpenOptions(params)

	v2 := map[string]string{
		"appid":     "wx842f10c69bb48e5b",
		"sub_appid": "wxa95825ecf5e840e6",
		"key":       "c5479de357f86743903ad11897b4e54c",
	}
	wechatEasySdk.NewOptions(v2)
}

func TestPartnerPay_Jsapi(t *testing.T) {

	options := wechatEasySdk.OpenOptions()
	bizContent := map[string]interface{}{
		"sp_appid":     options.SpAppid,
		"sp_mchid":     options.SpMchid,
		"out_trade_no": "1217752501201407033233368318",
		"sub_appid":    "wx1da941c68db4f659",
		"sub_mchid":    "1605553806",
		"description":  "Image形象店-深圳腾大-QQ公仔",
		"notify_url":   "https://console.mashangdian.cn/api/v1/wechat/auth_redirect",
		"amount": map[string]interface{}{
			"total":    1,
			"currency": "CNY",
		},
		"payer": map[string]interface{}{
			"sub_openid": "oZt0w5YP8a-3A4QwLJMiNLhAE9PE",
		},
	}

	data := new(PartnerPay).Jsapi(bizContent)
	fmt.Println("data", data)

}

func TestPartnerPay_Refunds(t *testing.T) {

	bizContent := map[string]interface{}{

		"sub_mchid":     "1605553806",
		"out_trade_no":  "T202105061605845925",
		"out_refund_no": "refund_1620272671",
		"amount": map[string]interface{}{
			"refund":   1680,
			"total":    1680,
			"currency": "CNY",
		},
	}

	data := new(PartnerPay).Refunds(bizContent)

	fmt.Println(data)

}

func TestPay_GetSignKey(t *testing.T) {

	bizContent := map[string]interface{}{
		"mch_id":    "1605269485",
		"nonce_str": "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
	}

	new(Pay).GetSignKey(bizContent)

	// c5479de357f86743903ad11897b4e54c

}

func TestPay_UnifiedOrder(t *testing.T) {

	bizContent := map[string]interface{}{
		"appid":            "wxd678efh567hg6787",
		"body":             "测试商品",
		"mch_id":           "1605269485",
		"nonce_str":        "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"notify_url":       "https://www.weixin.qq.com/wxpay/pay.php",
		"out_trade_no":     "sandbox_test_552",
		"spbill_create_ip": "192.168.10.10",
		"total_fee":        "552",
		"trade_type":       "JSAPI",
	}

	new(Pay).UnifiedOrder(bizContent)

}

func TestPay_OrderQuery(t *testing.T) {
	bizContent := map[string]interface{}{
		"appid":        "wxd678efh567hg6787",
		"mch_id":       "1605269485",
		"nonce_str":    "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"out_trade_no": "sandbox_test_552",
	}

	new(Pay).OrderQuery(bizContent)
}

func TestPay_Refund(t *testing.T) {
	bizContent := map[string]interface{}{
		"appid":         "wxd678efh567hg6787",
		"mch_id":        "1605269485",
		"nonce_str":     "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"out_refund_no": "TM6201505620352",
		"out_trade_no":  "sandbox_test_552",
		"refund_fee":    "552",
		"total_fee":     "552",
	}
	new(Pay).Refund(bizContent)
}

func TestPay_DownLoadBill(t *testing.T) {
	bizContent := map[string]interface{}{
		"appid":         "wxd678efh567hg6787",
		"mch_id":        "1605269485",
		"nonce_str":     "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"bill_date":     "20140603",
		"bill_type":     "ALL",
	}
	new(Pay).DownLoadBill(bizContent)
}

func TestPay_RefundQuery(t *testing.T) {
	bizContent := map[string]interface{}{
		"appid":         "wxd678efh567hg6787",
		"mch_id":        "1605269485",
		"nonce_str":     "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		"out_trade_no": "sandbox_test_552",
	}
	new(Pay).RefundQuery(bizContent)
}