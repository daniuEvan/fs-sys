package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Minio struct {
		BucketName string
		Endpoint   string
		AccessKey  string
		SecretKey  string
		UseSSL     bool
	}
}
