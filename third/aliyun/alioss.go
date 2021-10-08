/**
 * @Auth: Nuts
 * @Date: 2021/3/8 10:22 上午
 */
package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type AliossClient struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
}

func NewClient(accessKeyId, accessKeySecret, endpoint string) *AliossClient {
	return &AliossClient{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
	}
}

func (aliClient *AliossClient) UploadFile(localFile, bucketName, keyName string) error {
	// 创建OSSClient实例。
	client, err := oss.New(aliClient.Endpoint, aliClient.AccessKeyId, aliClient.AccessKeySecret)
	if err != nil {
		fmt.Println("Error1:", err)
		os.Exit(-1)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error2:", err)
		os.Exit(-1)
		return err

	}

	// 读取本地文件。
	fd, err := os.Open(localFile)
	if err != nil {
		fmt.Println("Error3:", err)
		os.Exit(-1)
		return err
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(keyName, fd)
	if err != nil {
		fmt.Println("Error4:", err)
		os.Exit(-1)
		return err

	}


	return nil

}
