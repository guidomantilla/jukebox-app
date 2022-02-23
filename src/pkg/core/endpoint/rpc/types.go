package rpc

import (
	"context"
	"fmt"
)

type DefaultGreetServiceServer struct {
}

func (server *DefaultGreetServiceServer) Greet(_ context.Context, request *GreetRequest) (*GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &GreetResponse{
		Result: result,
	}
	return res, nil
}
func (server *DefaultGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}
