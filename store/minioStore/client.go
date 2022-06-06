/**
 * @date: 2022/5/26
 * @desc:
 */

package minioStore

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

// NewMinioClient 创建客户端
func NewMinioClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	if minioClient != nil {
		return minioClient, nil
	}
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
