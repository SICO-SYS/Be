/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc"
	"log"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cfg/v2"
	"github.com/SiCo-Ops/dao/mongo"
	"github.com/SiCo-Ops/dao/redis"
)

var (
	config              cfg.ConfigItems
	configPool          = redis.NewPool()
	RPCServer           = grpc.NewServer()
	assetDB, assetDBErr = mongo.NewDial()
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

	configPool = redis.InitPool(config.RedisConfigHost, config.RedisConfigPort, config.RedisConfigAuth)
	configs, err := redis.Hgetall(configPool, "system.config")
	if err != nil {
		log.Fatalln(err)
	}
	cfg.Map2struct(configs, &config)

	assetDB, assetDBErr = mongo.InitDial(config.MongoAssetAddress, config.MongoAssetUsername, config.MongoAssetPassword)
	if assetDBErr != nil {
		log.Fatalln(assetDBErr)
	}

	pb.RegisterAssetServiceServer(RPCServer, &AssetService{})
	pb.RegisterTemplateServiceServer(RPCServer, &TemplateService{})

	if config.SentryBeStatus == "active" && config.SentryBeDSN != "" {
		raven.SetDSN(config.SentryBeDSN)
	}
}
