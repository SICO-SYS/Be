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
	"github.com/SiCo-Ops/dao/mongo"
	// "github.com/SiCo-Ops/public"
)

type TemplateService struct{}

func (t *TemplateService) CreateRPC(ctx context.Context, in *pb.AssetTemplateCall) (*pb.AssetTemplateBack, error) {
	c := mongo.CollectionTemplateName(in.Id)
	data := make(map[string]interface{})
	data["name"] = in.Name
	v := make(map[string]string)
	json.Unmarshal(in.Params, &v)
	param := []map[string]string{v}
	data["param"] = param
	err := mongo.TemplateEnsureIndexes(assetDB, in.Id)
	if err != nil {
		return &pb.AssetTemplateBack{Code: 203}, nil
	}
	err = mongo.Insert(assetDB, c, data)
	if err != nil {
		return &pb.AssetTemplateBack{Code: 203}, nil
	}
	return &pb.AssetTemplateBack{Code: 0}, nil
}
