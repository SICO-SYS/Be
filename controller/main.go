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
)

var (
	RPCServer = grpc.NewServer()
)

type AssetService struct{}

func init() {
	defer func() {
		recover()
	}()
	pb.RegisterAssetServiceServer(RPCServer, &AssetService{})

	if config.Sentry.Enable {
		raven.SetDSN(config.Sentry.DSN)
	}
}
