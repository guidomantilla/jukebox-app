package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"jukebox-app/examples/endpoints/grpc/rpc"
)

func main() {

	args := os.Args
	fmt.Println("Hello I'm a client")

	var err error
	var cc *grpc.ClientConn
	var mode string

	opts := grpc.WithInsecure()

	if len(args) == 1 && args[0] == "tls" {
		log.Fatal("missing parameters")
	}

	if len(args) == 1 && args[0] != "tls" {
		mode = args[0]
	}

	if len(args) == 2 && args[0] != "tls" {
		log.Fatal("tls parameter should be first")
	}

	if len(args) == 2 && args[0] == "tls" {
		mode = args[1]
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		cred, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(cred)
	}

	if cc, err = grpc.Dial("localhost:50051", opts); err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := rpc.NewGreetServiceClient(cc)
	doMode(mode, c)
}

func doMode(mode string, c rpc.GreetServiceClient) {
	switch mode {
	case "unary":
		doUnary(c)
	case "server-streaming":
		doServerStreaming(c)
	case "client-streaming":
		doClientStreaming(c)
	case "bidi-streaming":
		doBiDiStreaming(c)
	case "unary-deadline-ok":
		doUnaryWithDeadline(c, 5*time.Second)
	case "unary-deadline-bad":
		doUnaryWithDeadline(c, 1*time.Second)
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
		{
			Greeting: &rpc.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "Mark",
			},
		},
		{
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

func doBiDiStreaming(c rpc.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*rpc.GreetEveryoneRequest{
		{
			Greeting: &rpc.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "Mark",
			},
		},
		{
			Greeting: &rpc.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}

func doUnaryWithDeadline(c rpc.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do a UnaryWithDeadline RPC...")
	req := &rpc.GreetWithDeadlineRequest{
		Greeting: &rpc.Greeting{
			FirstName: "Stephane",
			LastName:  "Maarek",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {

		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
		return
	}
	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
