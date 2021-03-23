package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Himanshu/GoGrpc/greet/greetpb"
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
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Client creater: %f", c)
	doUnary(c)
	//doStream(c)
	//doClientStream(c)
	//doBiDiStreaming(c)
	//doUnaryWithDeadline(c, 5*time.Second)
	//doUnaryWithDeadline(c, 1*time.Second)
}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Himanshu",
			LastName:  "Jakhmola",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to call Greet RPC %v", err)

	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Stream RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Himanshu",
			LastName:  "jakhmola",
		},
	}
	reStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := reStream.Recv()
		if err == io.EOF {
			//Reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Stream RPC...")

	result := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Himanshu",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ron",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mikhel",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Albert",
			}},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Failed to call LongGreet: %v", err)
	}
	for _, req := range result {
		fmt.Printf("Sending Request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving response from LongGreet: %v", err)
	}
	fmt.Printf("Sending Request: %v\n", res)

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Client Stream RPC...")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Failed to call GreetEveryone: %v", err)
	}
	result := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Himanshu",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Ron",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mikhel",
			}},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Albert",
			}},
	}
	waitc := make(chan struct{})
	go func() {

		for _, req := range result {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reciving: %v\n", err)
			}
			fmt.Printf("Recived: %v\n", res)
		}
		close(waitc)
	}()
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do a UnaryWithDeadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Himanshu",
			LastName:  "Jakhmola",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit!! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Unable to call Greetwithdeadline RPC %v", err)
		}
		return

	}
	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
