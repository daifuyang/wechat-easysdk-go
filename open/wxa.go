/**
** @创建时间: 2021/5/3 6:32 下午
** @作者　　: return
** @描述　　:
 */
package open

import (
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
	"net/url"
)

type Wxa struct {
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 待授权小程序上传小程序代码
 * @Date 2021/5/3 18:45:40
 * @Param
 * @return
 **/
type ResultResponse struct {
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 设置服务器域名
 * @Date 2021/5/3 19:29:39
 * @Param
 * @return
 **/

type ModifyDomain struct {
	data.Response
	RequestDomain          []string `json:"requestdomain"`
	WsRequestDomain        []string `json:"wsrequestdomain"`
	UploadDomain           []string `json:"uploaddomain"`
	DownloadDomain         []string `json:"downloaddomain"`
	InvalidRequestDomain   []string `json:"invalid_requestdomain"`
	InvalidWsRequestDomain []string `json:"invalid_wsrequestdomain"`
	InvalidUploadDomain    []string `json:"invalid_uploaddomain"`
	InvalidDownloadDomain  []string `json:"invalid_downloaddomain"`
}

func (rest *Wxa) ModifyDomain(authorizerAccessToken string, bizContent map[string]interface{}) ModifyDomain {
	url := "https://api.weixin.qq.com/wxa/modify_domain?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(url, bizContent)
	result := ModifyDomain{}
	json.Unmarshal(data, &result)
	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 设置业务域名
 * @Date 2021/5/5 9:4:2
 * @Param
 * @return
 **/
type WebViewDomain struct {
	data.Response
	WebviewDomain []string `json:"webviewdomain"`
}

func (rest *Wxa) SetWebViewDomain(authorizerAccessToken string, bizContent map[string]interface{}) WebViewDomain {
	url := "https://api.weixin.qq.com/wxa/setwebviewdomain?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(url, bizContent)
	result := WebViewDomain{}
	json.Unmarshal(data, &result)
	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description
 * @Date 2021/5/5 9:20:44
 * @Param
 * @return
 **/

type templateList struct {
	CreateTime             int    `json:"create_time"`
	UserVersion            string `json:"user_version"`
	UserDesc               string `json:"user_desc"`
	TemplateId             int    `json:"template_id"`
	SourceMiniprogramAppid string `json:"source_miniprogram_appid"`
	SourceMiniprogram      string `json:"source_miniprogram"`
	Developer              string `json:"developer"`
}

type TemplateListResult struct {
	TemplateList []templateList `json:"template_list"`
	data.Response
}

func (rest *Wxa) GetTemplateList(componentAccessToken string) TemplateListResult {

	url := "https://api.weixin.qq.com/wxa/gettemplatelist?access_token=" + componentAccessToken
	data := util.GetJsonResult(url, nil)
	listResult := TemplateListResult{}
	json.Unmarshal(data, &listResult)
	return listResult
}

func (rest *Wxa) Commit(authorizerAccessToken string, bizContent map[string]interface{}) ResultResponse {

	url := "https://api.weixin.qq.com/wxa/commit?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(url, bizContent)
	fmt.Println(string(data))
	result := ResultResponse{}
	json.Unmarshal(data, &result)

	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取体验二维码
 * @Date 2021/5/3 18:46:38
 * @Param
 * @return
 **/

func (rest *Wxa) GetQrcode(authorizerAccessToken string, path string) (ResultResponse, []byte) {

	requestUrl := "https://api.weixin.qq.com/wxa/get_qrcode?access_token=" + authorizerAccessToken

	if path != "" {
		requestUrl += "&path=" + url.QueryEscape(path)
	}

	data := util.GetJsonResult(requestUrl, nil)

	result := ResultResponse{}
	json.Unmarshal(data, &result)

	return result, data

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 获取所有体验成员列表
 * @Date 2021/5/5 19:8:52
 * @Param
 * @return
 **/

type members struct {
	Userstr string `json:"userstr"`
}

type MemberAuth struct {
	Members []members `json:"members"`
	data.Response
}

func (rest *Wxa) MemberAuth(authorizerAccessToken string) MemberAuth {

	requestUrl := "https://api.weixin.qq.com/wxa/memberauth?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	bizContent["action"] = "get_experiencer"
	data := util.PostJsonResult(requestUrl, bizContent)
	result := MemberAuth{}
	json.Unmarshal(data, &result)
	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取小程序用户码
 * @Date 2021/5/5 20:7:55
 * @Param
 * @return
 **/
func (rest *Wxa) GetWxaCode(authorizerAccessToken string) (ResultResponse, []byte) {
	requestUrl := "https://api.weixin.qq.com/wxa/getwxacode?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	bizContent["path"] = "pages/index/index"
	data := util.PostJsonResult(requestUrl, bizContent)

	result := ResultResponse{}
	json.Unmarshal(data, &result)

	return result, data

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取微信的类目
 * @Date 2021/5/6 6:52:44
 * @Param
 * @return
 **/

type exterList struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type qualify struct {
	ExterList []exterList `json:"exter_list"`
	Remark    string      `json:"remark"`
}

type categories struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	Level              int      `json:"level"`
	Father             int      `json:"father"`
	Children           []int    `json:"children"`
	TypeList           []string `json:"type_list"`
	Qualify            qualify  `json:"qualify"`
	IsHidden           bool     `json:"is_hidden"`
	AvailableApiList   []string `json:"available_api_list"`
	Apis               []string `json:"apis"`
	Type               string   `json:"type"`
	AvailableForPlugin bool     `json:"available_for_plugin"`
	NeedReport         int      `json:"need_report"`
	CanUseCityserivce  int      `json:"can_use_cityserivce"`
	Scope              string   `json:"scope"`
	CanApply           bool     `json:"can_apply"`
	BizCannotUse       bool     `json:"biz_cannot_use"`
	BizSensitiveType   int      `json:"biz_sensitive_type"`
}

type categoriesList struct {
	Categories []categories `json:"categories,omitempty"`
}
type Categories struct {
	CategoriesList categoriesList `json:"categories_list,omitempty"`
	data.Response
}

func (rest *Wxa) GetAllCategories(authorizerAccessToken string) Categories {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/wxopen/getallcategories?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	data := util.GetJsonResult(requestUrl, bizContent)

	categoriesResult := Categories{}
	json.Unmarshal(data, &categoriesResult)

	return categoriesResult

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取审核时可填写的类目
 * @Date 2021/5/6 8:5:35
 * @Param
 * @return
 **/

func (rest *Wxa) GetCategories(authorizerAccessToken string) {
	requestUrl := "https://api.weixin.qq.com/wxa/get_category?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	data := util.GetJsonResult(requestUrl, bizContent)
	fmt.Println(string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 提交微信审核
 * @Date 2021/5/6 7:58:35
 * @Param
 * @return
 **/

type AuditResponse struct {
	Auditid int `json:"auditid"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 提交审核
 * @Date 2021/5/6 8:18:6
 * @Param
 * @return
 **/
func (rest *Wxa) SubmitAudit(authorizerAccessToken string, bizContent map[string]interface{}) AuditResponse {

	requestUrl := "https://api.weixin.qq.com/wxa/submit_audit?access_token=" + authorizerAccessToken
	data := util.PostJsonResult(requestUrl, bizContent)

	audit := AuditResponse{}
	json.Unmarshal(data, &audit)

	fmt.Println("audit", audit)
	return audit
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询指定发布审核单的审核状态
 * @Date 2021/5/6 8:18:13
 * @Param
 * @return
 **/

func (rest *Wxa) GetAuditStatus(authorizerAccessToken string, auditId int) {
	requestUrl := "https://api.weixin.qq.com/wxa/get_auditstatus?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	bizContent["auditid"] = auditId
	data := util.PostJsonResult(requestUrl, bizContent)
	fmt.Println("data", string(data))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 查询最新发布的审核状态
 * @Date 2021/5/6 12:24:38
 * @Param
 * @return
 **/

type LAuditResponse struct {
	Auditid    string `json:"auditid"`
	Status     int    `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"screen_shot"`
	data.Response
}

func (rest *Wxa) GetLatestAuditStatus(authorizerAccessToken string) LAuditResponse {
	requestUrl := "https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	data := util.GetJsonResult(requestUrl, bizContent)

	fmt.Println("string", string(data))

	lAuditResponse := LAuditResponse{}

	json.Unmarshal(data, &lAuditResponse)

	return lAuditResponse

}

// 发布已通过审核的小程序
func (rest *Wxa) Release(authorizerAccessToken string) data.Response {
	requestUrl := "https://api.weixin.qq.com/wxa/release?access_token=" + authorizerAccessToken
	bizContent := make(map[string]interface{}, 0)
	jsonData := util.PostJsonResult(requestUrl, bizContent)
	result := data.Response{}
	json.Unmarshal(jsonData, &result)
	return result
}

// 创建访问小程序的h5链接
type schemeResponse struct {
	OpenLink string `json:"openlink"`
	data.Response
}

func (rest *Wxa) GenerateScheme(authorizerAccessToken string, bizContent map[string]interface{}) schemeResponse {
	requestUrl := "https://api.weixin.qq.com/wxa/generatescheme?access_token=" + authorizerAccessToken
	jsonData := util.PostJsonResult(requestUrl, bizContent)
	result := schemeResponse{}
	json.Unmarshal(jsonData, &result)
	return result
}

type urlLinkResponse struct {
	UrlLink string `json:"url_link"`
	data.Response
}

func (rest *Wxa) GenerateUrlLink(authorizerAccessToken string, bizContent map[string]interface{}) urlLinkResponse {
	requestUrl := "https://api.weixin.qq.com/wxa/generate_urllink?access_token=" + authorizerAccessToken
	jsonData := util.PostJsonResult(requestUrl, bizContent)
	result := urlLinkResponse{}
	json.Unmarshal(jsonData, &result)
	return result
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取模板列表
 * @Date 2021/5/27 17:31:17
 * @Param
 * @return
 **/
func (rest *Wxa) GetPubTemplateTitles(authorizerAccessToken string) {
	requestUrl := "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=" + authorizerAccessToken + `&ids="632"&start=0&limit=1`

	jsonData := util.GetJsonResult(requestUrl, nil)

	fmt.Println("jsonData", string(jsonData))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取模板组合
 * @Date 2021/5/27 17:6:58
 * @Param
 * @return
 **/
func (rest *Wxa) GetPubTemplateKeywords(authorizerAccessToken string, tid string) {

	requestUrl := "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=" + authorizerAccessToken + "&tid=" + tid
	jsonData := util.GetJsonResult(requestUrl, nil)
	fmt.Println("jsonData", string(jsonData))

}

/**
 * @Author return <1140444693@qq.com>
 * @Description 组合模板到个人模板库
 * @Date 2021/5/27 17:32:40
 * @Param
 * @return
 **/

type AddTemplateResponse struct {
	data.Response
	PriTmplId string `json:"priTmplId"`
}

func (rest *Wxa) AddTemplate(authorizerAccessToken string, bizContent map[string]interface{}) AddTemplateResponse {

	requestUrl := "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token=" + authorizerAccessToken
	jsonData := util.PostJsonResult(requestUrl, bizContent)

	result := AddTemplateResponse{}
	json.Unmarshal(jsonData, &result)
	return result

}

/**
 * @Author return <1140444693@qq.com>
 * @Description // 获取个人模板库列表
 * @Date 2021/5/27 17:38:17
 * @Param
 * @return
 **/

func (rest *Wxa) GetTemplate(authorizerAccessToken string) {
	requestUrl := "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=" + authorizerAccessToken
	jsonData := util.GetJsonResult(requestUrl, nil)
	fmt.Println("jsonData", string(jsonData))
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 获取插件状态
 * @Date 2021/7/4 11:13:34
 * @Param
 * @return
 **/

type PluginList struct {
	AppId      string `json:"appid"`
	Status     int    `json:"status"`
	Nickname   string `json:"nickname"`
	HeadImgUrl string `json:"headimgurl"`
	Reason     string `json:"reason"`
}

type PluginListResponse struct {
	data.Response
	PluginList []PluginList `json:"plugin_list"`
}

func (rest *Wxa) GetPluginList(authorizerAccessToken string) PluginListResponse {
	requestUrl := "https://api.weixin.qq.com/wxa/plugin?access_token=" + authorizerAccessToken
	jsonData := util.PostJsonResult(requestUrl, map[string]interface{}{"action": "list"})

	pluginList := PluginListResponse{}
	json.Unmarshal(jsonData, &pluginList)

	return pluginList
}

type ApplyPluginResponse struct {
	data.Response
}

func (rest *Wxa) ApplyPlugin(authorizerAccessToken string, pluginAppId string) ApplyPluginResponse {

	requestUrl := "https://api.weixin.qq.com/wxa/plugin?access_token=" + authorizerAccessToken
	jsonData := util.PostJsonResult(requestUrl, map[string]interface{}{
		"plugin_appid": pluginAppId,
		"action":       "apply",
	})

	result := ApplyPluginResponse{}
	json.Unmarshal(jsonData, &result)

	return result

}
