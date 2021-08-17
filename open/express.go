/**
** @创建时间: 2021/5/8 8:39 下午
** @作者　　: return
** @描述　　: 即时配送
 */
package open

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/5/8 20:40:7
 * @Param 开通即时配送
 * @return
 **/

type Delivery struct{}

func (rest *Delivery) GetAll(authorizerAccessToken string) {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/delivery/getall?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	data := util.PostJsonResult(requestUrl, bizContent)
	fmt.Println(string(data))

}

func (rest *Delivery) Open(authorizerAccessToken string, deliveryId string) data.Response {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/open?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	bizContent["delivery_id"] = deliveryId
	resultResponse := util.PostJsonResult(requestUrl, bizContent)
	resultData := data.Response{}
	json.Unmarshal(resultResponse, &resultData)
	return resultData
}

// 第三方代商户发起绑定配送公司帐号的请求
func (rest *Delivery) Add(authorizerAccessToken string, deliveryId string) data.Response {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/shop/add?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	bizContent["delivery_id"] = deliveryId
	resultResponse := util.PostJsonResult(requestUrl, bizContent)
	resultData := data.Response{}
	json.Unmarshal(resultResponse, &resultData)
	return resultData

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 下配送预订单
 * @Date 2021/5/8 21:13:2
 * @Param
 * @return
 **/

type PreAddData struct {
	data.Response
	data.ImDeliveryResponse
	Fee              float64 `json:"fee"`
	DeliveryToken    string  `json:"delivery_token"`
	DispatchDuration int     `json:"dispatch_duration"`
}

func (rest *Delivery) PreAdd(authorizerAccessToken string, bizContent map[string]interface{}) PreAddData {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/order/pre_add?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)

	fmt.Println("预下单获取配送费", string(data))

	var preAddData = PreAddData{}
	json.Unmarshal(data, &preAddData)

	return preAddData

}

type OrderAddData struct {
	data.Response
	Fee              float64 `json:"fee"`
	WaybillId        string  `json:"waybill_id"`
	OrderStatus      int     `json:"order_status"`
	DispatchDuration int     `json:"dispatch_duration"`
}

func (rest *Delivery) OrderAdd(authorizerAccessToken string, bizContent map[string]interface{}) OrderAddData {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/order/add?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)

	fmt.Println("add data", string(data))

	var orderAddData = OrderAddData{}
	json.Unmarshal(data, &orderAddData)

	return orderAddData

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 测试更新订单状态
 * @Date 2021/5/8 22:45:21
 * @Param
 * @return
 **/
func (rest *Delivery) TestUpdateOrder(authorizerAccessToken string, bizContent map[string]interface{}) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/test_update_order?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 模拟配送公司更新配送单状态
 * @Date 2021/5/8 23:14:42
 * @Param
 * @return
 **/
func (rest *Delivery) RealMockUpdateOrder(authorizerAccessToken string, bizContent map[string]interface{}) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/realmock_update_order?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 拉取配送单信息
 * @Date 2021/5/8 23:1:26
 * @Param
 * @return
 **/
func (rest *Delivery) OrderGet(authorizerAccessToken string, bizContent map[string]interface{}) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/order/get?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 拉取已绑定账号
 * @Date 2021/6/6 8:39:18
 * @Param
 * @return
 **/

func (rest *Delivery) GetBindAccount(authorizerAccessToken string) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/express/local/business/shop/get?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, nil)
	fmt.Println(string(data))
}
