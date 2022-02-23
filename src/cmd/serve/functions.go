package serve

import (
	"fmt"
	"jukebox-app/src/endpoint/rpc"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func ExecuteCmdFn(_ *cobra.Command, _ []string) {
	fmt.Println("Hello World")

	var err error
	var lis net.Listener
	if lis, err = net.Listen("tcp", "0.0.0.0:50051"); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := make([]grpc.ServerOption, 0)
	server := grpc.NewServer(opts...)

	rpc.RegisterGreetServiceServer(server, &rpc.GreetServiceGrpcServer{})
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
