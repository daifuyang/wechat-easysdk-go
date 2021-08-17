/**
** @创建时间: 2021/5/3 6:32 下午
** @作者　　: return
** @描述　　:
 */
package open

import (
	"bytes"
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"github.com/gincmf/wechatEasySdk/util"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"
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

func TestComponent_Token(t *testing.T) {

	bizContent := map[string]interface{}{
		"component_appid":         "wxa95825ecf5e840e6",
		"component_appsecret":     "c733bbd5643a248fe19e7af3a837fcf1",
		"component_verify_ticket": "ticket@@@GBFWvbt4bCEVlEE7ky_1vTTtvug77Bydmy79LNtHOykHWtV9WcgB7jL7wH7DbZT_GOlzS7FLF8KXKSKjfVD1BA",
	}

	ak := new(Component).Token(bizContent)
	fmt.Println("ak", ak)

}

var ak = "46__jDMY3Y8ufKb0DMa7yVtoivGvxrDacj-b8vbcXGZI6_wF2L84ipCSaPkDPV2tHd6tF4NB5gEWnzLj4_zdcZFqwoFUE8v_s0a9c5SJXmr7q-GPif2TiZSUuuEyBt5VyRSwYnGz3GNlNK4FLP6WXXcALDPQQ"

func TestWxa_ModifyDomain(t *testing.T) {

	bizContent := map[string]interface{}{
		"action": "add",
		"requestdomain": []string{
			"https://console.mashangdian.cn",
		},
		"wsrequestdomain": []string{
			"wss://console.mashangdian.cn",
		},
		"uploaddomain": []string{
			"https://console.mashangdian.cn",
		},
		"downloaddomain": []string{
			"https://console.mashangdian.cn",
		},
	}

	data := new(Wxa).ModifyDomain(ak, bizContent)
	fmt.Println("data", data)

}

func TestWxa_SetWebViewDomain(t *testing.T) {

	bizContent := map[string]interface{}{
		"action": "add",
		"webviewdomain": []string{
			"https://console.mashangdian.cn",
		},
	}

	data := new(Wxa).SetWebViewDomain(ak, bizContent)
	fmt.Println("data", data)

}

func TestWxa_Commit(t *testing.T) {

	bizContent := map[string]interface{}{
		"template_id":  "16",
		"ext_json":     "{}",
		"user_version": "0.0.8",
		"user_desc":    "0.0.8beta体验版",
	}

	data := new(Wxa).Commit(ak, bizContent)

	fmt.Println("data", data)
}

func TestWxa_GetTemplateList(t *testing.T) {

	listResult := new(Wxa).GetTemplateList("44_NjKDQpOrDOnSXIqgzlXvWTCm87AMyjKqfIx2HhWr-sccFVc0_xfYn2DifLEIjFMEsZH074V1fI426r7DOCIUIggByfhmx86WnwitIq1zzXH1PeTvkRHBMBaku7zy7n9nhIwUCaJSmYfjf4-LTXBeAFAGKO")
	fmt.Println("listResult", listResult)

}

var cak = "46_2G9xY3vhbYaE-yejEdAkP_uJ5fWrBjlEXWvqo94Cib3oBLGphcwUY8HqbYDvHjJpAAls0Rfg2_gWbGM-H2dv-N3MDd8wqGcvNgWFIqwJ7XwZTO9Yy5MQAiRMorZTM8a3EEBuEaNpRzPMm5EIQRAgAJAMNW"

func TestComponent_FastRegisterWeapp(t *testing.T){

	bizContent := make(map[string]interface{}, 0)
	bizContent["name"] = "上海呵呵哒智能科技有限公司杨浦分公司"
	bizContent["code"] = "91310110MA1G98NF38"
	bizContent["code_type"] = 1
	bizContent["legal_persona_wechat"] = "chenrui33109"
	bizContent["legal_persona_name"] = "陈寿瑞"
	bizContent["component_phone"] = "17177723588"

	/*bizContent := make(map[string]interface{}, 0)
	bizContent["name"] = "码上云网络科技（温州）有限公司"
	bizContent["code"] = "91330301MA2HDMCN1P"
	bizContent["code_type"] = 1
	bizContent["legal_persona_wechat"] = "daifuyang123"
	bizContent["legal_persona_name"] = "戴富阳"
	bizContent["component_phone"] = "17177723588"*/

	result := new(Component).FastRegisterWeapp(cak, bizContent)
	fmt.Println("result", result)
}

func TestComponent_FastRegisterWeappSearch(t *testing.T) {
	bizContent := make(map[string]interface{}, 0)
	bizContent["name"] = "上海呵呵哒智能科技有限公司杨浦分公司"
	bizContent["legal_persona_wechat"] = "chenrui33109"
	bizContent["legal_persona_name"] = "陈寿瑞"
	result := new(Component).FastRegisterWeappSearch(cak, bizContent)
	fmt.Println("result", result)
}

func TestWxa_GetQrcode(t *testing.T) {
	response, jpg := new(Wxa).GetQrcode(ak, "")
	if response.Errcode == 0 {
		// 创建一个文件用于保存
		out, err := os.Create("test.jpg")
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// 然后将响应流和文件流对接起来
		_, err = io.Copy(out, ioutil.NopCloser(bytes.NewReader(jpg)))
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(response.Errmsg)
	}
}

func TestWxa_GetWxaCode(t *testing.T) {
	new(Wxa).GetWxaCode(ak)
}

func TestWxa_GetAllCategories(t *testing.T) {
	result := new(Wxa).GetAllCategories(ak)
	fmt.Println("result", result)
}

func TestWxa_GetCategories(t *testing.T) {
	new(Wxa).GetCategories(ak)
}

func TestWxa_SubmitAudit(t *testing.T) {
	bizContent := make(map[string]interface{}, 0)
	bizContent["version_desc"] = "帮助商家实快速实现数字化转型模板"
	bizContent["feedback_info"] = "仅作为演示账号，无实际经营用途。劳烦工作人员帮忙通过下"
	new(Wxa).SubmitAudit(ak, bizContent)
}

func TestWxa_GetAuditStatus(t *testing.T) {
	new(Wxa).GetAuditStatus(ak, 461122708)
}

func TestWxa_GetLatestAuditStatus(t *testing.T) {
	result := new(Wxa).GetLatestAuditStatus(ak)
	fmt.Println(result)
}

func TestLocalBusiness_GetAll(t *testing.T) {
	new(Delivery).GetAll(ak)
}

func TestLocalBusiness_Open(t *testing.T) {
	result :=new(Delivery).Open(ak, "DADA")
	fmt.Println(result)
}

func TestLocalBusiness_Add(t *testing.T) {
	result := new(Delivery).Add(ak, "DADA")
	fmt.Println(result)
}

func TestDelivery_GetBindAccount(t *testing.T) {
	new(Delivery).GetBindAccount(ak)
}

// 下预订单
func TestLocalBusiness_PreAdd(t *testing.T) {

	bizContent := make(map[string]interface{}, 0)
	bizContent["shopid"] = "test_shop_id"
	bizContent["shop_order_id"] = "DADA12345678"
	bizContent["shop_no"] = "001"

	dSign := bizContent["shopid"].(string) + bizContent["shop_order_id"].(string) + "test_app_secrect"
	bizContent["delivery_sign"] = util.GetSha1(dSign)
	bizContent["delivery_id"] = "TEST"
	bizContent["openid"] = "oZt0w5YP8a-3A4QwLJMiNLhAE9PE"

	bizContent["shop"] = map[string]interface{}{
		"wxa_path":    "/page/order/detail",
		"img_url":     "https://cdn.mashangdian.cn/tenant/2100695345/20210310/6d513d66f3ec23d9f95cf7e8c20cd5d0.jpeg!clipper",
		"goods_name":  "码上点外卖测试商品",
		"goods_count": 5,
	}

	bizContent["sender"] = map[string]interface{}{
		"name":            "戴富阳",
		"city":            "上海市",
		"address":         "宝山区",
		"address_detail":  "殷高西路",
		"phone":           "17177723588",
		"lng":             121.485014,
		"lat":             31.319462,
		"coordinate_type": 0,
	}

	bizContent["receiver"] = map[string]interface{}{
		"name":            "测试",
		"city":            "北京市",
		"address":         "北京市海淀区宝盛东路与宝盛南路交叉路口往北约100米(奥北科技园东侧)\"",
		"address_detail":  "私语茶舍",
		"phone":           "15161178722",
		"lng":             116.384264,
		"lat":             40.039247,
		"coordinate_type": 0,
	}

	bizContent["cargo"] = map[string]interface{}{
		"goods_value":        30,
		"goods_weight":       1,
		"cargo_first_class":  "美食夜宵",
		"cargo_second_class": "快餐/地方菜",
	}

	data :=  new(Delivery).PreAdd(ak, bizContent)
	fmt.Println(data)
}

// 下正式单
func TestLocalBusiness_OrderAdd(t *testing.T) {

	bizContent := make(map[string]interface{}, 0)

	bizContent["delivery_token"] = "test_delivery_token"

	bizContent["shopid"] = "test_shop_id"
	bizContent["shop_order_id"] = "DADA12345678"
	bizContent["shop_no"] = "001"

	dSign := bizContent["shopid"].(string) + bizContent["shop_order_id"].(string) + "d80400f91e156f63b38886e616d84590"
	bizContent["delivery_sign"] = util.GetSha1(dSign)
	bizContent["delivery_id"] = "TEST"
	bizContent["openid"] = "oZt0w5YP8a-3A4QwLJMiNLhAE9PE"

	bizContent["shop"] = map[string]interface{}{
		"wxa_path":    "/page/order/detail",
		"img_url":     "https://cdn.mashangdian.cn/tenant/2100695345/20210310/6d513d66f3ec23d9f95cf7e8c20cd5d0.jpeg!clipper",
		"goods_name":  "码上点外卖测试商品",
		"goods_count": 5,
	}

	bizContent["sender"] = map[string]interface{}{
		"name":            "戴富阳",
		"city":            "北京市",
		"address":         "北京北京市海淀区国泰大厦",
		"address_detail":  "国泰大厦",
		"phone":           "15161178723",
		"lng":             116.377441,
		"lat":             40.032368,
		"coordinate_type": 0,
	}

	bizContent["receiver"] = map[string]interface{}{
		"name":            "测试",
		"city":            "北京市",
		"address":         "北京市海淀区宝盛东路与宝盛南路交叉路口往北约100米(奥北科技园东侧)\"",
		"address_detail":  "私语茶舍",
		"phone":           "15161178722",
		"lng":             116.384264,
		"lat":             40.039247,
		"coordinate_type": 0,
	}

	bizContent["cargo"] = map[string]interface{}{
		"goods_value":        30,
		"cargo_first_class":  "美食夜宵",
		"cargo_second_class": "快餐/地方菜",
	}

	bizContent["order_info"] = map[string]interface{}{
		"order_time": time.Now().Unix(),
	}

	data := new(Delivery).OrderAdd(ak, bizContent)

	fmt.Println(data)
}

// 更新订单状态
func TestDelivery_TestUpdateOrder(t *testing.T) {

	bizContent := make(map[string]interface{}, 0)
	bizContent["shopid"] = "test_shop_id"
	bizContent["shop_order_id"] = "DADA202106161911272974"
	bizContent["action_time"] = time.Now().Unix()
	bizContent["order_status"] = 202
	dSign := bizContent["shopid"].(string) + bizContent["shop_order_id"].(string) + "test_app_secrect"
	bizContent["delivery_sign"] = util.GetSha1(dSign)
	new(Delivery).TestUpdateOrder(ak, bizContent)

}

func TestDelivery_RealMockUpdateOrder(t *testing.T) {

	bizContent := make(map[string]interface{}, 0)
	bizContent["shopid"] = "1534713176"
	bizContent["shop_order_id"] = "SFTC2000000000"
	bizContent["action_time"] = time.Now().Unix()
	bizContent["order_status"] = 302
	dSign := bizContent["shopid"].(string) + bizContent["shop_order_id"].(string) + "d80400f91e156f63b38886e616d84590"
	bizContent["delivery_sign"] = util.GetSha1(dSign)

	new(Delivery).RealMockUpdateOrder(ak, bizContent)
}

// 获取配送单信息
func TestDelivery_OrderGet(t *testing.T) {

	bizContent := make(map[string]interface{}, 0)
	bizContent["shopid"] = "test_shop_id"
	bizContent["shop_order_id"] = "test_shop_order_id"
	bizContent["shop_no"] = "123123"
	dSign := bizContent["shopid"].(string) + bizContent["shop_order_id"].(string) + "test_app_secrect"
	bizContent["delivery_sign"] = util.GetSha1(dSign)
	new(Delivery).OrderGet(ak, bizContent)

}

func TestSubscribe_Send(t *testing.T) {

	keywords := map[string]interface{}{
		"name1": map[string]string{
			"value": "测试门店",
		},
		"character_string4": map[string]string{
			"value": "test1234",
		},
		"amount3": map[string]string{
			"value": "取餐号：1234",
		},
		"date2": map[string]string{
			"value": "2021-05-09",
		},
		"thing5": map[string]string{
			"value": "备注",
		},
	}

	bizContent := map[string]interface{}{
		"touser":            "oZt0w5YP8a-3A4QwLJMiNLhAE9PE",
		"template_id":       "6rhlDbE3jLNQSqZQ-vKRx2vzuTZj8-W4-vnMPV0hs4s",
		"page":              "pages/order/detail?id=10",
		"miniprogram_state": "developer",
		"data":              keywords,
	}

	new(Subscribe).Send(ak, bizContent)
}

func TestWxa_GetPubTemplateTitles(t *testing.T) {
	new(Wxa).GetPubTemplateTitles(ak)
}

func TestWxa_GetPubTemplateKeywords(t *testing.T) {
	new(Wxa).GetPubTemplateKeywords(ak, "677")
}

func TestWxa_AddTemplate(t *testing.T) {
	addResponse := new(Wxa).AddTemplate(ak, map[string]interface{}{
		"tid":       "264",
		"kidList":   []string{"45", "37", "12", "4", "26"},
		"sceneDesc": "点餐下单成功通知",
	})

	if addResponse.Errcode == 0 {

	}
}

func TestWxa_GetTemplate(t *testing.T) {
	new(Wxa).GetTemplate(ak)
}

func TestWxa_GetPluginList(t *testing.T) {
	pluginList := new(Wxa).GetPluginList(ak)
	fmt.Println("pluginList",pluginList)
}

func TestWxa_ApplyPlugin(t *testing.T) {
	response := new(Wxa).ApplyPlugin(ak,"wxaae6519cee98d824")
	fmt.Println("response",response)
}

