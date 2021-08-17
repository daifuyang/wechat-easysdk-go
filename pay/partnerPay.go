/**
** @创建时间: 2021/4/22 1:38 下午
** @作者　　: return
** @描述　　:
 */
package pay

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

type PartnerPay struct {
	PrepayId string `json:"prepay_id"`
	data.PartnerResponse
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 服务商收款
 * @Date 2021/4/22 13:37:0
 * @Param
 * @return
 **/
func (rest *PartnerPay) Jsapi(bizContent map[string]interface{}) PartnerPay {
	url := "/v3/pay/partner/transactions/jsapi"
	data, err := util.GetPartnerJsonResult(url, bizContent)

	if err != nil {
		fmt.Println("err",err.Error())
	}

	fmt.Println("data",string(data))

	pay := PartnerPay{}
	json.Unmarshal(data, &pay)
	return pay
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 服务商退款
 * @Date 2021/5/6 10:17:44
 * @Param
 * @return
 **/

type amount struct {
	Currency         string `json:"currency"`
	DiscountRefund   int    `json:"discount_refund"`
	PayerRefund      int    `json:"payer_refund"`
	PayerTotal       int    `json:"payer_total"`
	Refund           int    `json:"refund"`
	SettlementRefund int    `json:"settlement_refund"`
	SettlementTotal  int    `json:"settlement_total"`
	Total            int    `json:"total"`
}

type RefundsResponse struct {
	Amount amount `json:"amount,omitempty"`
	data.PartnerResponse
}

func (rest *PartnerPay) Refunds(bizContent map[string]interface{}) RefundsResponse {
	url := "/v3/refund/domestic/refunds"
	data, _ := util.GetPartnerJsonResult(url, bizContent)
	refundsResponse := RefundsResponse{}
	json.Unmarshal(data, &refundsResponse)
	return refundsResponse
}


/**
 * @Author return <1140444693@qq.com>
 * @Description 
 * @Date 2021/5/13 23:27:54
 * @Param
 * @return
 **/