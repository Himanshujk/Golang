package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Himanshu/GoGrpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello i am Client!!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("Client creater: %f", c)
	//doUnary(c)
	//doPrime(c)
	doSquareRoot(c)
}
func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  34,
		SecondNumber: 34,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to call Greet RPC %v", err)

	}
	log.Printf("Response from Greet: %v", res.SumResponse)
}
func doPrime(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Prime number Decomposition RPC..")
	req := &calculatorpb.PrimeNumberRequest{
		Number: 120,
	}

	stream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to call PrimeNumberDecomposition RPC %v", err)

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in streaming: %v\n", err)
		}
		fmt.Println(res.GetPrimeNumber())
	}

}
func doSquareRoot(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Implementing SquareRoot unary RPC..!\n")
	//correct call
	doErrorcall(c, 45)
	//error call
	doErrorcall(c, -1)
}
func doErrorcall(c calculatorpb.CalculatorServiceClient, n float32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			//actual error from grpc
			fmt.Println(resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("Sended a negetive number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v\n", err)
			return
		}
	}
	fmt.Printf("Result of squre root of %v: %v\n", n, res.GetResult())
}
