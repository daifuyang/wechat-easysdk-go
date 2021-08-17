/**
** @创建时间: 2021/4/27 4:55 下午
** @作者　　: return
** @描述　　:
 */
package merchant

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gincmf/wechatEasySdk/data"
	"github.com/gincmf/wechatEasySdk/util"
	"io"
	"mime/multipart"
	"os"
	fp "path/filepath"
)

type Media struct {
	MediaId string `json:"media_id"`
	data.Response
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 商户上传
 * @Date 2021/4/27 16:58:36
 * @Param
 * @return
 **/
func (rest *Media) Upload(filepath string) (media Media,err error) {
	url := "/v3/merchant/media/upload"

	// 文件上传
	header := make(map[string]string, 0)

	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println("加载文件失败", err)
	}

	var meta struct {
		Filename string `json:"filename"`
		Sha256   string `json:"sha256"`
	}

	filename := fp.Base(file.Name())
	meta.Filename = filename

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println(err)
	}

	sha256New := sha256.New()
	sha256New.Write(buf.Bytes())

	fileSha256 := hex.EncodeToString(sha256New.Sum([]byte("")))
	meta.Sha256 = fileSha256

	metaStr, _ := json.Marshal(&meta)
	bodyWrite.WriteField("meta", string(metaStr))

	fileWrite, err := bodyWrite.CreateFormFile("file", filename)

	_, err = io.Copy(fileWrite, buf)
	if err != nil {
		fmt.Println("io Copy error", err)
	}

	header["Content-Type"] = bodyWrite.FormDataContentType()
	bodyWrite.Close()

	result, err := util.GetUploadResult(url, "POST", metaStr, bodyBuf, header)

	if err != nil {
		return media,err
	}

	json.Unmarshal(result,&media)

	return media,nil

}
