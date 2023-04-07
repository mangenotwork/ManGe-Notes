package conn

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//测试阿里云对象存储连接
func TestConnAliOSS(endPoint, accessKeyID, ossSecret, bucketName string) (bool, error) {
	client, err := oss.New(endPoint, accessKeyID, ossSecret)
	if err != nil {
		return false, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return false, err
	}
	log.Println(bucket)
	return true, nil
}
