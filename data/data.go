/**
** @创建时间: 2021/4/20 10:36 下午
** @作者　　: return
** @描述　　:
 */
package data

type Response struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type PartnerResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ImDeliveryResponse struct {
	Resultcode int `json:"resultcode,omitempty"`
	Resultmsg  string `json:"resultmsg,omitempty"`
}
