/**
** @创建时间: 2020/9/7 9:36 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"sort"
	"strings"
)

func Sign(content []string) (sign string) {

	contentStr := strings.Join(content, "\n")+"\n"

	h := sha256.New()
	h.Write([]byte(contentStr))

	options := wechatEasySdk.OpenOptions()

	block := []byte(options.PrivateKey)

	blocks, _ := pem.Decode(block)
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		fmt.Println("err", err.Error())
		return sign
	}

	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}

// 对参数签名，获取签名参数
func V2Sign(params map[string]interface{}) (sign string, encode string) {

	//ksort 对参数进行排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 对参数进行序列化
	pStr := make([]string, 0)
	//拼接
	for _, k := range keys {
		v := []byte(params[k].(string))
		v = bytes.TrimSpace(v)
		if string(v) != "" {
			key := params[k].(string)
			pStr = append(pStr, k+"="+key)
		}
	}

	// 序列化结果
	encode = strings.Join(pStr, "&")
	options := wechatEasySdk.Options()

	encode += "&key=" + options.Key

	sign = GetMd5(encode)

	sign = strings.ToUpper(sign)

	return sign, encode
}
