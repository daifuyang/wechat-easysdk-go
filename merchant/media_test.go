/**
** @创建时间: 2021/4/27 5:41 下午
** @作者　　: return
** @描述　　:
 */
package merchant

import (
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
		"gatewayHost": "https://api.mch.weixin.qq.com",
		"appCertPath": "../pem",
	}

	wechatEasySdk.NewOpenOptions(params)
}

func TestMerchant_Upload(t *testing.T) {
	new(Media).Upload("/Users/return/Downloads/logo.png")
}
