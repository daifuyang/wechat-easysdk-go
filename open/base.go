/**
** @创建时间: 2021/4/22 9:55 下午
** @作者　　: return
** @描述　　: 基础信息设置
 */

package open

import (
	"encoding/json"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

// 功能介绍信息
type signatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

// 头像信息
type headImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

// 名称信息
type nicknameInfo struct {
	Nickname        string `json:"nickname"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type Base struct {
	AppId             string        `json:"app_id"`
	AccountType       int           `json:"account_type"`
	PrincipalType     int           `json:"principal_type"`
	PrincipalName     string        `json:"principal_name"`
	SignatureInfo     signatureInfo `json:"signature_info"`
	HeadImageInfo     headImageInfo `json:"head_image_info"`
	NicknameInfo      nicknameInfo  `json:"nickname_info"`
	RegisteredCountry int           `json:"registered_country"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 调用本API 可以获取小程序的基本信息
 * @Date 2021/4/22 22:1:31
 * @Param
 * @return
 **/
func (rest *Base) BaseInfo(authorizerAccessToken string) Base {
	url := "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(url, nil)
	result := Base{}
	json.Unmarshal(data, &result)
	return result
}
