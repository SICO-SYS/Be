/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"encoding/json"
	"golang.org/x/net/context"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cloud-go-sdk/qcloud"
	"github.com/SiCo-Ops/dao/mongo"
	"github.com/SiCo-Ops/public"
)

type AssetService struct{}

func (a *AssetService) SynchronizeRPC(ctx context.Context, in *pb.AssetSynchronizeCall) (*pb.AssetMsgBack, error) {
	switch in.Cloud {
	case "qcloud":
		v := &qcloudSDK.CVM{}
		json.Unmarshal(in.Data, v)
		collection := mongo.CollectionAssetCloudName(in.Cloud, in.Id)
		mongo.Remove(mongo.AssetConn, collection)
		for _, cloudResource := range v.Response.InstanceSet {
			mongo.Insert(mongo.AssetConn, cloudResource, collection)
		}
		return &pb.AssetMsgBack{Code: 0, Msg: public.Int642String(v.Response.TotalCount)}, nil
	default:
		return &pb.AssetMsgBack{Code: -1}
	}
	return &pb.AssetMsgBack{Code: -1}
}

func (a *AssetService) CustomRPC(ctx context.Context, in *pb.AssetCustomCall) (*pb.AssetMsgBack, error) {
	return &pb.AssetMsgBack{Code: -1}, nil
}
