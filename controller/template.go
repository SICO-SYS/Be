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

	"github.com/SiCo-DevOps/Pb"
	"github.com/SiCo-DevOps/dao/mongo"
	// "github.com/SiCo-DevOps/public"
)

func (a *Asset) AssetTemplate(ctx context.Context, in *pb.Asset_Req) (*pb.Asset_Res, error) {
	c := "asset.template." + in.Id
	data := make(map[string]string)
	data["name"] = in.Name
	for key, value := range in.Param {
		data[key] = value
	}
	ok := mongo.Mgo_Insert(mongo.MgoAssetConn, data, c)
	if ok {
		return &pb.Asset_Res{0}, nil
	}
	return &pb.Asset_Res{1}, nil
}

// func init() {
// }
