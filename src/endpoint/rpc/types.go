package rpc

import (
	"context"
	"fmt"
	"strconv"
	"time"
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
		res := &GreetManytimesResponse{
			Result: result,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (server *GreetServiceGrpcServer) mustEmbedUnimplementedGreetServiceServer() {}
