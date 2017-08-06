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

func (a *Asset) AssetTemplate(ctx context.Context, in *pb.Asset_Req) (*pb.Asset_Res, error) {
	c := "asset." + in.Id + ".template"
	data := make(map[string]interface{})
	data["name"] = in.Name
	param := []map[string]string{in.Param}
	data["param"] = param
	mongo.Asset_ensureIndexes(in.Id)
	ok := mongo.Mgo_Insert(mongo.MgoAssetConn, data, c)
	if ok {
		return &pb.Asset_Res{0}, nil
	}
	return &pb.Asset_Res{1}, nil
}

func (a *Asset) AssetStorein(ctx context.Context, in *pb.Asset_CloudReq) (*pb.Asset_Res, error) {
	return &pb.Asset_Res{}, nil
}
