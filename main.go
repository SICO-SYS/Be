/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package main

import (
	"net"

	"github.com/SiCo-DevOps/Be/controller"
	"github.com/SiCo-DevOps/cfg"
)

func Run() {
	lis, _ := net.Listen("tcp", cfg.Config.Rpc.Be)
	controller.S.Serve(lis)
}

func main() {
	Run()
}
