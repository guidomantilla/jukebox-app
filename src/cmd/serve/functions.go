package serve

import (
	"fmt"
	"jukebox-app/src/endpoint/rpc"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {
	fmt.Println("Hello World")

	var err error
	var lis net.Listener
	if lis, err = net.Listen("tcp", "0.0.0.0:50051"); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := make([]grpc.ServerOption, 0)
	if len(args) != 0 && args[0] == "tls" {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		var cred credentials.TransportCredentials
		if cred, err = credentials.NewServerTLSFromFile(certFile, keyFile); err != nil {
			log.Fatalf("Failed loading certificates: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(cred))
	}

	server := grpc.NewServer(opts...)
	rpc.RegisterGreetServiceServer(server, &rpc.GreetServiceGrpcServer{})
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
