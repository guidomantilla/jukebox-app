package test

import (
	"context"
	"fmt"
	"jukebox-app/src/endpoint/rpc"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func ExecuteCmdFn(_ *cobra.Command, args []string) {
	fmt.Println("Hello I'm a client")

	var err error
	var cc *grpc.ClientConn

	opts := grpc.WithInsecure()
	if cc, err = grpc.Dial("localhost:50051", opts); err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := rpc.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c rpc.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &rpc.GreetRequest{
		Greeting: &rpc.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
