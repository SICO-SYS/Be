/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"google.golang.org/grpc"

	"github.com/SiCo-DevOps/Pb"
)

var (
	S   = grpc.NewServer()
	err error
)

type Asset struct{}

func init() {
	defer func() {
		recover()
	}()
	pb.RegisterAseetServer(S, &Asset{})
}
