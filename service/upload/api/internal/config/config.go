package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Uploader struct {
		AsyncUpload      bool
		FileTempLocation string
		BucketName       string
	}
	Minio struct {
		Endpoint  string
		AccessKey string
		SecretKey string
		UseSSL    bool
	}
}
