/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"golang.org/x/net/context"
	// "io/ioutil"
	// "net/http"
	// "sort"
	// "strings"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/dao/mongo"
	// "github.com/SiCo-Ops/public"
)

type TemplateService struct{}

func (t *TemplateService) CreateRPC(ctx context.Context, in *pb.AssetTemplateCall) (*pb.AssetMsgBack, error) {
	c := "template." + in.Id
	data := make(map[string]interface{})
	data["name"] = in.Name
	param := []map[string]string{in.Param}
	data["param"] = param
	mongo.TemplateEnsureIndexes(in.Id)
	ok := mongo.Insert(mongo.AssetConn, data, c)
	if ok {
		return &pb.AssetMsgBack{Code: 0}, nil
	}
	return &pb.AssetMsgBack{Code: 1}, nil
}
