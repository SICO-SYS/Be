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
)

type AssetService struct{}

func (a *AssetService) SynchronizeRPC(ctx context.Context, in *pb.AssetSynchronizeCall) (*pb.AssetSynchronizeBack, error) {
	collection := mongo.CollectionAssetCloudName(in.Cloud, in.Id)
	err := mongo.Remove(assetDB, collection, nil)
	if err != nil {
		return &pb.AssetSynchronizeBack{Code: 203}, nil
	}
	switch in.Cloud {
	case "qcloud":
		switch in.Service {
		case "cvm":
			v := &qcloudSDK.CVM{}
			json.Unmarshal(in.Data, v)
			for _, cloudResource := range v.Response.InstanceSet {
				mongo.Insert(assetDB, collection, cloudResource)
			}
			return &pb.AssetSynchronizeBack{Code: 0, TotalCount: v.Response.TotalCount}, nil
		default:
			return &pb.AssetSynchronizeBack{Code: 3000}, nil
		}
	case "aliyun":
		switch in.Service {
		case "ecs":
			v := &aliyunSDK.ECS{}
			json.Unmarshal(in.Data, v)
			for _, cloudResource := range v.Instances.Instance {
				mongo.Insert(assetDB, collection, cloudResource)
			}
			return &pb.AssetSynchronizeBack{Code: 0, TotalCount: v.TotalCount}, nil
		default:
			return &pb.AssetSynchronizeBack{Code: 3000}, nil
		}
	case "aws":
		switch in.Service {
		case "ec2":
			v := &awsSDK.EC2DescribeInstances{}
			xml.Unmarshal(in.Data, v)
			for _, reservation := range v.ReservationSet {
				for _, instance := range reservation.InstancesSet {
					mongo.Insert(assetDB, collection, instance)
				}
			}
			if v.NextToken != "" {
				return &pb.AssetSynchronizeBack{Code: 0, NextToken: v.NextToken}, nil
			}
			return &pb.AssetSynchronizeBack{Code: 0}, nil
		default:
			return &pb.AssetSynchronizeBack{Code: 3000}, nil
		}
	default:
		return &pb.AssetSynchronizeBack{Code: 3000}, nil
	}
}

func (a *AssetService) CustomRPC(ctx context.Context, in *pb.AssetCustomizeCall) (*pb.AssetCustomizeBack, error) {
	return &pb.AssetCustomizeBack{Code: 0}, nil
}
