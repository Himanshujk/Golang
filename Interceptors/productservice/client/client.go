package main

import (
	"context"
	"fmt"

	productservice "github.com/Himanshu/Interceptors/productservice/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(
		"localhost:1111",
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	productServ := productservice.NewProductServiceClient(conn)

	response1, err1 := productServ.FindAll(context.Background(), &productservice.FindAllRequest{})
	if err1 != nil {
		fmt.Println(err1)
	} else {
		products := response1.Products
		fmt.Println("Product List")
		for _, product := range products {
			fmt.Println("id: ", product.Id)
			fmt.Println("name: ", product.Name)
			fmt.Println("price: ", product.Price)
			fmt.Println("quantity: ", product.Quantity)
			fmt.Println("status: ", product.Status)
			fmt.Println("========================")
		}
	}

	response2, err2 := productServ.Search(context.Background(), &productservice.SearchRequest{Keyword: "vi"})
	if err2 != nil {
		fmt.Println(err2)
	} else {
		products := response2.Products
		fmt.Println("Search Products")
		for _, product := range products {
			fmt.Println("id: ", product.Id)
			fmt.Println("name: ", product.Name)
			fmt.Println("price: ", product.Price)
			fmt.Println("quantity: ", product.Quantity)
			fmt.Println("status: ", product.Status)
			fmt.Println("========================")
		}
	}

}
