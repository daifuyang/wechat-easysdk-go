/**
** @创建时间: 2021/4/22 10:10 下午
** @作者　　: return
** @描述　　:
 */
package open

import (
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"strings"
	"testing"
)

func init() {

	params := map[string]string{
		"spAppid":               "wx842f10c69bb48e5b",
		"spMchid":               "1605269485",
		"appId":                 "wxa95825ecf5e840e6",
		"appSecret":             "c733bbd5643a248fe19e7af3a837fcf1",
		"aesKey":                "codecloud2021codecloud2021codecloud20212021=",
		"appCertPath":           "../pem",
		"componentVerifyTicket": "ticket@@@QLXHNjyGZM4SSTW8IqQBSrDHieq-FbMXz1xsV_zyR6lU0odmHyCCWZnYL7GjPZwoyKp9zmDzkS_hFB5yC22UTg",
	}

	wechatEasySdk.NewOpenOptions(params)
}

func TestComponent_AuthorizerToken(t *testing.T) {

	options := wechatEasySdk.OpenOptions()
	bizContent := map[string]interface{}{
		"component_appid":         options.AppId,
		"component_appsecret":     options.AppSecret,
		"component_verify_ticket": options.ComponentVerifyTicket,
	}

	fmt.Println("bizContent", bizContent)

	comToken := new(Component).Token(bizContent)
	fmt.Println("comToken", comToken)
	accessToken := strings.TrimSpace(comToken.AccessToken)

	bizContent = map[string]interface{}{
		"component_appid":          options.AppId,
		"authorizer_appid":         "wx1da941c68db4f659",
		"authorizer_refresh_token": "refreshtoken@@@EXLcrs8Z6yCGW81wfH-V7D35eA_5fSjM9kIsyZ95B3Q",
	}
	aToken := new(Component).AuthorizerToken(accessToken, bizContent)
	fmt.Println("aToken",aToken.AuthorizerAccessToken)
}

func TestBase_BaseInfo(t *testing.T) {
	new(Base).BaseInfo("44_UNR7cnQWyL4kyzkwkXNqIQAaFDIS_Owkoz1N9oajo-EtwfG41acri4xAIuHa65QJ9mrAvDrgaQqDO77htlyv8g8m1PE_kKrbiSr5Kvlhcq_i0a-40j17h_V7HQNFaw3hf5wks4Zd4NA5aClVBLEdAHDOYR")
}
