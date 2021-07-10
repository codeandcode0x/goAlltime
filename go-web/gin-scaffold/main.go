//main
package main

import (
	"gin-scaffold/route"
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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
		// RpcServer()
		log.Println("rpc server stop ....")
		return nil
	})

	endless.ListenAndServe(":8080", r)
}

// rpc server
// func RpcServer() {
// 	log.Println("rpc server runing .....")
// 	lis, err := net.Listen("tcp", ":20153")
// 	if err != nil {
// 		fmt.Printf("failed to listen: %v", err)
// 		return
// 	}

// 	s := grpc.NewServer()

// 	jobPro.RegisterCheckTeamResBlockServer(s, &jobRpc.JobRpcServer{})
// 	reflection.Register(s)
// 	err = s.Serve(lis)

// 	if err != nil {
// 		fmt.Printf("failed to serve: %v", err)
// 		return
// 	}
// }
