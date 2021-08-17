/**
** @创建时间: 2021/1/4 11:19 下午
** @作者　　: return
** @描述　　:
 */
package pay

import (
	"fmt"
	"github.com/gincmf/wechatEasySdk/util"
)

type Pay struct {
}

type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	Quantity         int    `json:"quantity"`
	UnitPrice        int    `json:"unit_price"`
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 商户收款
 * @Date 2021/4/22 13:37:0
 * @Param
 * @return
 **/
func (rest *Pay) MicroPay(bizContent map[string]interface{}) {
	url := "https://api.mch.weixin.qq.com/pay/micropay"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 测试用例
 * @Date 2021/5/13 23:23:13
 * @Param
 * @return
 **/
func (rest *Pay) GetSignKey(bizContent map[string]interface{}) {
	url := "/sandboxnew/pay/getsignkey"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 测试下单
 * @Date 2021/5/13 23:50:51
 * @Param
 * @return
 **/
func (rest *Pay) UnifiedOrder(bizContent map[string]interface{}) {
	url := "/sandboxnew/pay/unifiedorder"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 测试查询订单
 * @Date 2021/5/14 0:27:18
 * @Param
 * @return
 **/
func (rest *Pay) OrderQuery(bizContent map[string]interface{}) {
	url := "/sandboxnew/pay/orderquery"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <11退款40444693@qq.com>
 * @Description //申请
 * @Date 2021/5/14 0:30:1
 * @Param
 * @return
 **/
func (rest *Pay) Refund(bizContent map[string]interface{}) {
	url := "/sandboxnew/pay/refund"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 交易对账单下载
 * @Date 2021/5/14 0:39:46
 * @Param
 * @return
 **/
func (rest *Pay) DownLoadBill(bizContent map[string]interface{}) {
	url := "/sandboxnew/pay/downloadbill"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询退款
 * @Date 2021/5/14 0:48:33
 * @Param
 * @return
 **/

func (rest *Pay) RefundQuery(bizContent map[string]interface{}) {

	url := "/sandboxnew/pay/refundquery"
	data := util.GetResult(url, bizContent)
	fmt.Println(string(data))
}
