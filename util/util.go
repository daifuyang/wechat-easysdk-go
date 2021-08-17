/**
** @创建时间: 2020/9/7 9:46 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/gincmf/wechatEasySdk"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 封装请求库
func Request(method string, url string, body io.Reader, header map[string]string) (int, []byte) {
	client := &http.Client{}
	switch method {
	case "get", "GET":
		method = "GET"
	case "post", "POST":
		method = "POST"
	case "put", "PUT":
		method = "PUT"
	case "delete", "DELETE":
		method = "POST"
	}
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("http错误", err)
	}

	r.Header.Add("Host", "")
	r.Header.Add("Accept", "*/*")
	r.Header.Add("Connection", "keep-alive")
	r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	r.Header.Add("Content-Length", "0")
	r.Header.Add("Cache-Control", "no-cache")
	for k, v := range header {
		r.Header.Add(k, v)
	}

	response, err := client.Do(r)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	var data []byte = nil

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			data = append(data, buf...)
		}
	default:
		data, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "image/jpeg" {
		index := bytes.IndexByte(data, 0)
		if index > 0 {
			data = data[:index]
		}
	}

	return response.StatusCode, data
}

// 获取证书序列号
func GetSerialNo() (serialNo string, err error) {

	options := wechatEasySdk.OpenOptions()

	block := []byte(options.PublicKey)

	blocks, _ := pem.Decode(block)

	x509Cert, err := x509.ParseCertificate(blocks.Bytes)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}

	serialNo = x509Cert.SerialNumber.Text(16)
	return serialNo, nil
}

// 自动获取序列话签名操作
func EncodeParams(paramsMap map[string]interface{}) string {
	// 获取签名
	sign, _ := V2Sign(paramsMap)

	paramsMap["sign"] = sign // 追加参数

	// 获取提交的参数列表
	xmlBytes, _ := xml.Marshal(Xml(paramsMap))

	return string(xmlBytes)
}

// 自动获取请求头
func getBearerToken(url string, methods string, jsonStr []byte) (authorization string, err error) {
	// 获取提交的参数列表
	options := wechatEasySdk.OpenOptions()

	mchid := options.SpMchid

	unix := time.Now().Unix()

	timestamp := strconv.FormatInt(unix, 10)

	nonceStr := GetMd5(timestamp)
	bodyJson := string(jsonStr)

	methods = strings.ToUpper(methods)

	encryptData := []string{
		methods,
		url,
		timestamp,
		nonceStr,
		bodyJson,
	}

	signature := Sign(encryptData)
	serialNo, err := GetSerialNo()
	if err != nil {
		return "", err
	}

	access := []string{
		"mchid=\"" + mchid + "\"",
		"nonce_str=\"" + nonceStr + "\"",
		"signature=\"" + signature + "\"",
		"timestamp=\"" + timestamp + "\"",
		"serial_no=\"" + serialNo + "\"",
	}

	accessToken := strings.Join(access, ",")

	schema := "WECHATPAY2-SHA256-RSA2048"

	authorization = schema + " " + accessToken

	return authorization, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description MD5
 * @Date 2021/1/8 23:29:27
 * @Param
 * @return
 **/

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetResult(url string, paramsMap map[string]interface{}) []byte {
	return request(url, paramsMap, nil, nil)
}

// 服务商微信支付json请求
func jsonRequest(url string, method string, paramsMap map[string]interface{}, body io.Reader, header map[string]string) []byte {

	// 获取提交的参数列表

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(paramsMap); err != nil {}
	jsonStr := bf.Bytes()

	if header == nil {
		header = make(map[string]string, 0)
	}

	header = map[string]string{
		"Content-Type": "application/json;charset=UTF-8",
	}

	_, data := Request(method, url, bytes.NewBuffer(jsonStr), header)
	return data
}

func GetJsonResult(url string, paramsMap map[string]interface{}) []byte {
	return jsonRequest(url, "GET", paramsMap, nil, nil)
}

func PostJsonResult(url string, paramsMap map[string]interface{}) []byte {
	return jsonRequest(url, "POST", paramsMap, nil, nil)
}

func request(url string, paramsMap map[string]interface{}, body io.Reader, header map[string]string) []byte {

	xml := EncodeParams(paramsMap)


	header = map[string]string{
		"Content-Type": "application/ssml+xml",
	}

	options := wechatEasySdk.OpenOptions()
	url = options.GatewayHost + url

	_, data := Request("POST", url, bytes.NewBuffer([]byte(xml)), header)
	return data
}

/*微信支付服务商请求*/
func GetPartnerJsonResult(url string, paramsMap map[string]interface{}) (data []byte, err error) {

	data, err = partnerJsonRequest(url, "POST", paramsMap, nil, nil)

	if err != nil {
		return data, err
	}

	return data, nil
}

/*微信支付服务商请求*/
func partnerJsonRequest(url string, methods string, paramsMap map[string]interface{}, body io.Reader, header map[string]string) (data []byte, err error) {

	// 获取提交的参数列表
	options := wechatEasySdk.OpenOptions()

	var jsonStr []byte
	if paramsMap != nil {
		jsonStr, _ = json.Marshal(paramsMap)
	}

	accessToken, err := getBearerToken(url, methods, jsonStr)

	if err != nil {
		return data, err
	}

	if header == nil {
		header = make(map[string]string, 0)
	}

	header["Authorization"] = accessToken
	header["Content-Type"] = "application/json;charset=UTF-8"
	header["Accept"] = "application/json;charset=UTF-8"

	url = options.GatewayHost + url
	_, data = Request(methods, url, bytes.NewBuffer(jsonStr), header)
	return data, nil
}

func PartnerGetJsonRequest(url string, paramsMap map[string]interface{}, body io.Reader, header map[string]string) (data []byte, err error) {
	return partnerJsonRequest(url, "GET", paramsMap, body, header)
}

func PartnerPostJsonRequest(url string, paramsMap map[string]interface{}, body io.Reader, header map[string]string) (data []byte, err error) {
	return partnerJsonRequest(url, "POST", paramsMap, body, header)
}

// 上传文件
func GetUploadResult(url string, method string, jsonStr []byte, body io.Reader, header map[string]string) (data []byte, err error) {

	options := wechatEasySdk.OpenOptions()

	accessToken, err := getBearerToken(url, method, jsonStr)

	if err != nil {
		return data, err
	}

	header["Authorization"] = accessToken

	url = options.GatewayHost + url
	_, data = Request(method, url, body, header)
	return data, nil
}

// AES加密
//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个,然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS7填充的反向操作，删除填充的字符串
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	paddLen := int(origData[length-1])

	end := length - paddLen

	text := origData[:length]
	if end > 0 {
		text = origData[:end]
	}

	return text
}

//aes cbc加密操作
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)

	iv := make([]byte, blockSize)

	//采用AES加密方法中的CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//aes cbc解密操作
func AesDeCrypt(cypted []byte, key []byte, appointIv ...string) ([]byte, error) {

	key, _ = base64.StdEncoding.DecodeString(string(key))

	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	iv := make([]byte, blockSize)

	if len(appointIv) > 0 {
		iv, _ = base64.StdEncoding.DecodeString(appointIv[0])
	}

	//采用AES加密方法中的CBC加密模式 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cypted))
	//这个函数还可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData = PKCS7UnPadding(origData)
	return origData, err
}

// 敏感信息加密
func EncryptCiphertext(str string, v3 bool) (cipherText string, err error) {

	options := wechatEasySdk.OpenOptions()
	secretMessage := []byte(str)
	rng := rand.Reader

	block := []byte(options.PublicKey)

	if v3 {
		block = []byte(options.V3PublicKey)
	}

	blocks, _ := pem.Decode(block)
	pk, err := x509.ParseCertificate(blocks.Bytes)
	if err != nil {
		fmt.Println("x509 err", err.Error())
		return "", err
	}
	pub := pk.PublicKey

	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, pub.(*rsa.PublicKey), secretMessage, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "", err
	}

	cipherText = base64.StdEncoding.EncodeToString(cipherData)
	return cipherText, nil
}

// 敏感数据解密
func DeCryptCiphertext(str string) (cipherText string, err error) {

	options := wechatEasySdk.OpenOptions()
	secretMessage := []byte(str)
	rng := rand.Reader

	block := []byte(options.PrivateKey)

	blocks, _ := pem.Decode(block)
	pub, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		fmt.Println("x509 err", err.Error())
		return "", err
	}
	cipherData, err := rsa.DecryptOAEP(sha1.New(), rng, pub.(*rsa.PrivateKey), secretMessage, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "", err
	}

	return string(cipherData), nil
}

// 敏感数据支付平台解密
func AesDecrypt256Gcm(key, associatedData, nonce, ciphertext string) (string, error) {

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("NewCipher err", err)
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("aesgcm err", err)
		return "", err
	}

	plaintext, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		fmt.Println("plaintext err", err)
		return "", err
	}

	return string(plaintext), nil

}
