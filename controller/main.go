/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
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
}
