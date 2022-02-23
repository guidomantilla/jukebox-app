package test

import (
	"context"
	"fmt"
	"io"
	"jukebox-app/src/endpoint/rpc"
	"log"
	"time"

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

	mode := args[0]
	switch mode {
	case "unary":
		doUnary(c)
	case "server-streaming":
		doServerStreaming(c)
	case "client-streaming":
		doClientStreaming(c)
	}

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

func doServerStreaming(c rpc.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &rpc.GreetManyTimesRequest{
		Greeting: &rpc.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c rpc.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*rpc.LongGreetRequest{
		&rpc.LongGreetRequest{
			Greeting: &rpc.Greeting{
				FirstName: "Stephane",
			},
		},
		&rpc.LongGreetRequest{
			Greeting: &rpc.Greeting{
				FirstName: "John",
			},
		},
		&rpc.LongGreetRequest{
			Greeting: &rpc.Greeting{
				FirstName: "Lucy",
			},
		},
		&rpc.LongGreetRequest{
			Greeting: &rpc.Greeting{
				FirstName: "Mark",
			},
		},
		&rpc.LongGreetRequest{
			Greeting: &rpc.Greeting{
				FirstName: "Piper",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}
