package main

import (
	"fmt"
	"log"
	"net"

	utils "github.com/usefss/cab-please/identity/utils"
	"github.com/usefss/cab-please/surge/stubs"
	"google.golang.org/grpc"
)

func main() {
	port := utils.GetEnv("PORT")
	fmt.Println("Server is listening on port: " + port)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := GRPCServer{}
	grpcServer := grpc.NewServer()
	stubs.RegisterSurgeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
