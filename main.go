/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package main

import (
	"net"

	"github.com/SiCo-Ops/Be/controller"
	"github.com/SiCo-Ops/cfg"
)

func Run() {
	lis, _ := net.Listen("tcp", cfg.Config.RPCPort.Be)
	controller.RPCServer.Serve(lis)
}

func main() {
	Run()
}
