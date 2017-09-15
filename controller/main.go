/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cfg/v2"
	"github.com/SiCo-Ops/dao/mongo"
	"github.com/SiCo-Ops/dao/redis"
)

var (
	config              cfg.ConfigItems
	configPool          = redis.Pool("", "", "")
	RPCServer           = grpc.NewServer()
	assetDB, assetDBErr = mongo.Dial("", "", "")
)

func ServePort() string {
	return config.RpcBePort
}

func init() {
	defer func() {
		recover()
	}()
	data := cfg.ReadLocalFile()

	if data != nil {
		cfg.Unmarshal(data, &config)
	}

	configPool = redis.Pool(config.RedisConfigHost, config.RedisConfigPort, config.RedisConfigAuth)
	configs, _ := redis.Hgetall(configPool, "system.config")
	cfg.Map2struct(configs, &config)

	assetDB, assetDBErr = mongo.Dial(config.MongoAssetAddress, config.MongoAssetUsername, config.MongoAssetPassword)

	pb.RegisterAssetServiceServer(RPCServer, &AssetService{})
	pb.RegisterTemplateServiceServer(RPCServer, &TemplateService{})

	if config.SentryHeStatus == "active" && config.SentryHeDSN != "" {
		raven.SetDSN(config.SentryHeDSN)
	}
}
