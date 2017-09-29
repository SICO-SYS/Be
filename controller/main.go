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
	"github.com/SiCo-Ops/cfg"
	"github.com/SiCo-Ops/dao/mongo"
)

const (
	configPath string = "config.json"
)

var (
	config              cfg.ConfigItems
	RPCServer           = grpc.NewServer()
	assetDB, assetDBErr = mongo.NewDial()
)

func ServePort() string {
	return config.RpcBePort
}

func init() {
	data, err := cfg.ReadFilePath(configPath)
	if err != nil {
		data = cfg.ReadConfigServer()
		if data == nil {
			log.Fatalln("config.json not exist and configserver was down")
		}
	}
	cfg.Unmarshal(data, &config)

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
