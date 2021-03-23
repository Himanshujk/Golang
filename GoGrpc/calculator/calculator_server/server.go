package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/Himanshu/GoGrpc/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Sum(ctc context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum Service RPC is invoked with %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResponse: sum,
	}
	return res, nil
}
func (*server) PrimeNumber(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberServer) error {
	fmt.Printf("Prime Decomposition RPC is invoked with %v\n", req)
	num := req.GetNumber()
	divisior := int64(2)
	for num > 1 {
		if num%divisior == 0 {
			stream.Send(&calculatorpb.PrimeNumberResponse{
				PrimeNumber: divisior,
			})
			num = num / divisior
		} else {
			divisior++
			fmt.Printf("Divisior has increased to :%v\n", divisior)
		}
	}
	return nil
}
func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("SquareRoot Service RPC is invoked with %v\n", req)
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recieved negative number: %v\n", number),
		)
	}
	return &calculatorpb.SquareRootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil

}

func main() {
	fmt.Println("Calculator Server!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Listener Error: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%v", err)
	}

}
