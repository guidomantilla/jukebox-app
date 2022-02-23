package rpc

import (
	"context"
	"fmt"
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
func (server *GreetServiceGrpcServer) mustEmbedUnimplementedGreetServiceServer() {}
