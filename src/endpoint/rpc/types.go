package rpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreetServiceGrpcServer struct {
}

func (server *GreetServiceGrpcServer) Greet(_ context.Context, request *GreetRequest) (*GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &GreetResponse{
		Result: result,
	}
	return res, nil
}

func (server *GreetServiceGrpcServer) GreetManyTimes(request *GreetManyTimesRequest, stream GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &GreetManyTimesResponse{
			Result: result,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		fmt.Println(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (server *GreetServiceGrpcServer) LongGreet(stream GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (server *GreetServiceGrpcServer) GreetEveryone(stream GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Printf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func (server *GreetServiceGrpcServer) GreetWithDeadline(ctx context.Context, request *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", request)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &GreetWithDeadlineResponse{
		Result: result,
	}
	return res, nil
}

func (server *GreetServiceGrpcServer) mustEmbedUnimplementedGreetServiceServer() {}
