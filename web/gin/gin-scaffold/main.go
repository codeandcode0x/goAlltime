//main
package main

import (
	"gin-scaffold/route"
	"log"
	"net"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	g errgroup.Group
)

// @title gin scaffold
// @version 1.0
// @description  gin scaffold
// @termsOfService
// @Host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	route.DefinitionRoute(r)
	// rpc goroutine
	g.Go(func() error {
		RpcServer()
		log.Println("rpc server stop ....")
		return nil
	})

	//get server port
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = viper.GetString("SERVER_PORT")
	}

	log.Println("http server runing :" + serverPort)

	//end server
	endless.ListenAndServe(":"+serverPort, r)
}

// rpc server
func RpcServer() {
	// get gRPC port
	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		gRPCPort = viper.GetString("GRPC_PORT")
	}
	// net listen
	lis, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	log.Println("rpc server runing :" + gRPCPort)
	// new server
	s := grpc.NewServer()
	// register your gRPC func such as :
	// rpcName.RegisterYourFuncServer(s, &rpcName.YourRpcServer{})
	reflection.Register(s)
	err = s.Serve(lis)
	// return error
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return
	}
}
