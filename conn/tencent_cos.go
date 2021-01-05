package conn

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 请求示例,使用永久密钥
func TestConnTencentCOS(cos_url, secretid, secretkey string) (bool, error) {
	u, _ := url.Parse(cos_url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretid,
			SecretKey: secretkey,
		},
	})
	log.Println(client)
	_, _, err := client.Service.Get(context.Background())
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
