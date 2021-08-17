/**
** @创建时间: 2021/4/20 5:37 下午
** @作者　　: return
** @描述　　:
 */
package open

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
)

type Component struct{}

type AccessToken struct {
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	data.Response
}

type PreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
	data.Response
}

type funcscopeCategory struct {
	Id int `json:"id"`
}

type funcInfo struct {
	FuncscopeCategory funcscopeCategory `json:"funcscope_category"`
}

type authorizationInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfo               []funcInfo `json:"func_info"`
}

type AuthorizationResult struct {
	AuthorizationInfo authorizationInfo `json:"authorization_info"`
	data.Response
}

type AuthorizerAccessToken struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/4/20 22:13:19
 * @Param
 * @return
 **/
func (rest *Component) Token(bizContent map[string]interface{}) AccessToken {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	data := util.PostJsonResult(url, bizContent)
	accessToken := AccessToken{}
	json.Unmarshal(data, &accessToken)
	return accessToken
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 令牌（component_access_token）是第三方平台接口的调用凭据。令牌的获取是有限制的，每个令牌的有效期为 2 小时，请自行做好令牌的管理，在令牌快过期时（比如1小时50分），重新调用接口获取。
 * @Date 2021/4/20 17:41:26
 * @Param
 * @return
 **/

func (rest *Component) PreAuthCode(accessToken string, bizContent map[string]interface{}) PreAuthCode {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + accessToken
	data := util.PostJsonResult(url, bizContent)
	preAuthCode := PreAuthCode{}
	json.Unmarshal(data, &preAuthCode)
	return preAuthCode
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 授权
 * @Date 2021/4/21 1:10:7
 * @Param
 * @return
 **/
func (rest *Component) QueryAuth(accessToken string, bizContent map[string]interface{}) AuthorizationResult {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + accessToken
	data := util.PostJsonResult(url, bizContent)
	fmt.Println("data", string(data))
	authorizationResult := AuthorizationResult{}
	json.Unmarshal(data, &authorizationResult)
	return authorizationResult
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 在公众号/小程序接口调用令牌（authorizer_access_token）失效时，可以使用刷新令牌（authorizer_refresh_token）获取新的接口调用令牌。
 * @Date 2021/4/22 22:36:29
 * @Param
 * @return
 **/
func (rest *Component) AuthorizerToken(componentAccessToken string, bizContent map[string]interface{}) AuthorizerAccessToken {
	url := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=" + componentAccessToken
	data := util.PostJsonResult(url, bizContent)
	authorizerAccessToken := AuthorizerAccessToken{}
	json.Unmarshal(data, &authorizerAccessToken)
	return authorizerAccessToken
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 第三方平台开发者的服务器使用登录凭证（code）以及第三方平台的 component_access_token 可以代替小程序实现登录功能 获取 session_key 和 openid。其中 session_key 是对用户数据进行加密签名的密钥。为了自身应用安全，session_key 不应该在网络上传输。
 * @Date 2021/4/21 10:17:21
 * @Param
 * @return
 **/
type Code2Session struct {
	AppId                string `json:"app_id,omitempty"`
	Secret               string `json:"secret,omitempty"`
	JsCode               string `json:"js_code,omitempty"`
	ComponentAppid       string `json:"component_appid,omitempty"`
	ComponentAccessToken string `json:"component_access_token,omitempty"`
	GrantType            string `json:"grant_type,omitempty"`
}

type Code2SessionResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid,omitempty"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 自调用接口
 * @Date 2021/7/16 17:17:34
 * @Param
 * @return
 **/
func (rest *Component) SecretCode2session(params Code2Session) Code2SessionResult {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + params.AppId + "&secret=" + params.Secret + "&js_code=" + params.JsCode + "&grant_type=authorization_code"
	data := util.PostJsonResult(url, nil)
	code2SessionResult := Code2SessionResult{}
	json.Unmarshal(data, &code2SessionResult)
	return code2SessionResult
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 第三方平台开发者的服务器使用登录凭证（code）以及第三方平台的 component_access_token 可以代替小程序实现登录功能 获取 session_key 和 openid。其中 session_key 是对用户数据进行加密签名的密钥。为了自身应用安全，session_key 不应该在网络上传输。
 * @Date 2021/4/21 10:25:28
 * @Param
 * @return
 **/
func (rest *Component) Code2session(params Code2Session) Code2SessionResult {
	url := "https://api.weixin.qq.com/sns/component/jscode2session?appid=" + params.AppId + "&js_code=" + params.JsCode + "&grant_type=authorization_code&component_appid=" + params.ComponentAppid + "&component_access_token=" + params.ComponentAccessToken
	data := util.PostJsonResult(url, nil)
	code2SessionResult := Code2SessionResult{}
	json.Unmarshal(data, &code2SessionResult)
	return code2SessionResult
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 快速创建小程序
 * @Date 2021/5/6 15:48:43
 * @Param
 * @return
 **/

func (rest *Component) FastRegisterWeapp(ComponentAccessToken string, bizContent map[string]interface{}) (result data.Response) {
	url := "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=create&component_access_token=" + ComponentAccessToken
	data := util.PostJsonResult(url, bizContent)

	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询申请单
 * @Date 2021/5/12 15:54:29
 * @Param
 * @return
 **/

func (rest *Component) FastRegisterWeappSearch(ComponentAccessToken string, bizContent map[string]interface{}) (result data.Response) {
	url := "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp?action=search&component_access_token=" + ComponentAccessToken
	data := util.PostJsonResult(url, bizContent)

	json.Unmarshal(data, &result)

	return result

}
