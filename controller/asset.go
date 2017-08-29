/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cloud-go-sdk/qcloud"
)

type AssetService struct{}

func (a *AssetService) SynchronizeRPC(ctx context.Context, in *pb.AssetSynchronizeCall) (*pb.AssetMsgBack, error) {
	v := &qcloudSDK.CVM{}
	json.Unmarshal(in.Data, v)
	for index, cloudResource := range v.Response.InstanceSet {
		fmt.Println(index)
		fmt.Println(cloudResource.InstanceId)
	}
	return &pb.AssetMsgBack{Code: 1}, nil
}

func (a *AssetService) CustomRPC(ctx context.Context, in *pb.AssetCustomCall) (*pb.AssetMsgBack, error) {
	return &pb.AssetMsgBack{Code: 0}, nil
}
