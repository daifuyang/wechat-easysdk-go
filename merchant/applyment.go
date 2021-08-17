/**
** @创建时间: 2021/4/29 9:57 下午
** @作者　　: return
** @描述　　: 服务商（银行、支付机构、电商平台不可用）使用该接口提交商家资料，帮助商家入驻成为微信支付的特约商户。
 */
package merchant

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

type Applyment struct {
	ApplymentId int `json:"applyment_id"`
	data.PartnerResponse
}

type showApplyment struct {
	SubMchid          string              `json:"sub_mchid"`
	ApplymentId       int                 `json:"applyment_id"`
	ApplymentState    string              `json:"applyment_state"`
	ApplymentStateMsg string              `json:"applyment_state_msg"`
	AuditDetail       []map[string]string `json:"audit_detail"`
	BusinessCode      string              `json:"business_code"`
	SignUrl           string              `json:"sign_url"`
	data.PartnerResponse
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 提交申请单
 * @Date 2021/5/1 13:32:2
 * @Param
 * @return
 **/
func (rest *Applyment) Applyment(bizContent map[string]interface{}) Applyment {
	url := "/v3/applyment4sub/applyment/"
	header := make(map[string]string, 0)

	options := wechatEasySdk.OpenOptions()
	header["Wechatpay-Serial"] = options.WechatpaySerial

	data, _ := util.PartnerPostJsonRequest(url, bizContent, nil, header)

	fmt.Println("data", string(data))

	applyment := Applyment{}

	json.Unmarshal(data, &applyment)

	return applyment
}

// 查询申请单
/**
 * @Author return <1140444693@qq.com>
 * @Description 查询申请单
 * @Date 2021/5/1 13:32:13
 * @Param
 * @return
 **/
func (rest *Applyment) State(applymentId string) showApplyment {
	url := "/v3/applyment4sub/applyment/applyment_id/" + applymentId
	header := make(map[string]string, 0)

	options := wechatEasySdk.OpenOptions()
	header["Wechatpay-Serial"] = options.WechatpaySerial

	data, _ := util.PartnerGetJsonRequest(url, nil, nil, header)

	sa := showApplyment{}
	json.Unmarshal(data, &sa)

	return sa
}
