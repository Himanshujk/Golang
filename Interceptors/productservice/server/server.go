package main

import (
	"fmt"
	"net"

	"github.com/Himanshu/Interceptors/productservice/handlers"
	interceptors "github.com/Himanshu/Interceptors/productservice/interceptor"
	productservice "github.com/Himanshu/Interceptors/productservice/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	s := handlers.ProductServiceServer{}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptors.DateLogInterceptor,
				interceptors.MethodLogInterceptor,
			),
		),
	)

	productservice.RegisterProductServiceServer(server, &s)

	if err := server.Serve(lis); err != nil {
		fmt.Println(err)
	}

}
