package handlers

import (
	"context"
	"strings"

	productservice "github.com/Himanshu/Interceptors/productservice/proto"
)

type ProductServiceServer struct {
}

var products []*productservice.Product

func init() {
	products = []*productservice.Product{
		{
			Id:       "p01",
			Name:     "tivi 1",
			Price:    4.5,
			Quantity: 4,
			Status:   true,
		},
		{
			Id:       "p02",
			Name:     "tivi 2",
			Price:    22,
			Quantity: 3,
			Status:   false,
		},
		{
			Id:       "p03",
			Name:     "laptop 3",
			Price:    27,
			Quantity: 3,
			Status:   true,
		},
	}
}

func (*ProductServiceServer) FindAll(ctx context.Context, in *productservice.FindAllRequest) (*productservice.FindAllResponse, error) {
	return &productservice.FindAllResponse{Products: products}, nil
}

func (*ProductServiceServer) Search(ctx context.Context, in *productservice.SearchRequest) (*productservice.SearchResponse, error) {
	var result []*productservice.Product
	for _, product := range products {
		if strings.Contains(product.Name, in.Keyword) {
			result = append(result, product)
		}
	}
	return &productservice.SearchResponse{Products: result}, nil
}
