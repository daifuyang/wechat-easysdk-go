/**
** @创建时间: 2021/5/9 3:14 下午
** @作者　　: return
** @描述　　:
 */
package open

import (
	"encoding/json"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

type Subscribe struct {
}

func (rest *Subscribe) Send(authorizerAccessToken string, bizContent map[string]interface{}) (response data.Response) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)
	json.Unmarshal(data, &response)
	return response

}
