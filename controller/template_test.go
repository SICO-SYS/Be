/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"encoding/json"
	"golang.org/x/net/context"
	"testing"

	"github.com/SiCo-Ops/Pb"
	"github.com/SiCo-Ops/public"
)

func Test_CreateRPC(t *testing.T) {
	test := &TemplateService{}
	param := map[string]string{"key1": "value1"}
	params, _ := json.Marshal(param)
	in := &pb.AssetTemplateCall{Id: "1234567890abcdef", Name: "test", Params: params}
	res, err := test.CreateRPC(context.Background(), in)
	if err != nil {
		t.Error(err)
	}
	if res.Code != 0 {
		t.Error(res.Code)
	}
}

func Benchmark_CreateRPC(b *testing.B) {
	test := &TemplateService{}
	param := map[string]string{"key1": "value1"}
	params, _ := json.Marshal(param)
	for i := 0; i < b.N; i++ {
		in := &pb.AssetTemplateCall{Id: public.GenerateHexString(), Name: "test", Params: params}
		res, err := test.CreateRPC(context.Background(), in)
		if err != nil {
			b.Error(err)
		}
		if res.Code != 0 {
			b.Error(res.Code)
		}
	}
}
