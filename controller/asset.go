/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"golang.org/x/net/context"

	"github.com/SiCo-Ops/Pb"
)

func (a *AssetService) AddAssetRPC(ctx context.Context, in *pb.AssetCloudCall) (*pb.AssetMsgBack, error) {
	return &pb.AssetMsgBack{Code: 0}, nil
}
