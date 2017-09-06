/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"encoding/json"
	"encoding/xml"
	"golang.org/x/net/context"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/cloud-go-sdk/aliyun"
	"github.com/SiCo-Ops/cloud-go-sdk/aws"
	"github.com/SiCo-Ops/cloud-go-sdk/qcloud"
	"github.com/SiCo-Ops/dao/mongo"
	"github.com/SiCo-Ops/public"
)

type AssetService struct{}

func (a *AssetService) SynchronizeRPC(ctx context.Context, in *pb.AssetSynchronizeCall) (*pb.AssetMsgBack, error) {
	collection := mongo.CollectionAssetCloudName(in.Cloud, in.Id)
	mongo.Remove(mongo.AssetConn, collection)
	switch in.Cloud {
	case "qcloud":
		switch in.Service {
		case "cvm":
			v := &qcloudSDK.CVM{}
			json.Unmarshal(in.Data, v)
			for _, cloudResource := range v.Response.InstanceSet {
				mongo.Insert(mongo.AssetConn, cloudResource, collection)
			}
			return &pb.AssetMsgBack{Code: 0, Msg: public.Int642String(v.Response.TotalCount)}, nil
		default:
			return &pb.AssetMsgBack{Code: -1}, nil
		}
	case "aliyun":
		switch in.Service {
		case "ecs":
			v := &aliyunSDK.ECS{}
			json.Unmarshal(in.Data, v)
			for _, cloudResource := range v.Instances.Instance {
				mongo.Insert(mongo.AssetConn, cloudResource, collection)
			}
			return &pb.AssetMsgBack{Code: 0, Msg: public.Int642String(v.TotalCount)}, nil
		default:
			return &pb.AssetMsgBack{Code: -1}, nil
		}
	case "aws":
		switch in.Service {
		case "ec2":
			v := &awsSDK.EC2DescribeInstances{}
			xml.Unmarshal(in.Data, v)
			for _, reservation := range v.ReservationSet {
				for _, instance := range reservation.InstancesSet {
					mongo.Insert(mongo.AssetConn, instance, collection)
				}
			}
			return &pb.AssetMsgBack{Code: -1}, nil
		}
	default:
		return &pb.AssetMsgBack{Code: -1}, nil
	}
	return &pb.AssetMsgBack{Code: -1}, nil
}

func (a *AssetService) CustomRPC(ctx context.Context, in *pb.AssetCustomCall) (*pb.AssetMsgBack, error) {
	return &pb.AssetMsgBack{Code: -1}, nil
}
