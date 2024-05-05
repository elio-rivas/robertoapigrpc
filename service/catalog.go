package service

import (
	"context"
	pb "github.com/elio-rivas/robertoapigrpc/proto"
)

type CatalogService struct {
	pb.UnimplementedCatalogServiceServer
}

func NewCatalogService() *CatalogService {
	return &CatalogService{}
}

func (s *CatalogService) GetCatalog(ctx context.Context, req *pb.GetCatalogRequest) (*pb.GetCatalogResponse, error) {
	// Implement your method here
	return &pb.GetCatalogResponse{}, nil
}

func (s *CatalogService) CreateCatalog(ctx context.Context, req *pb.CreateCatalogItemRequest) (*pb.CatalogItem, error) {
	// Implement your method here
	return &pb.CatalogItem{}, nil
}

func (s *CatalogService) UpdateCatalog(ctx context.Context, req *pb.UpdateCatalogItemRequest) (*pb.CatalogItem, error) {
	// Implement your method here
	return &pb.CatalogItem{}, nil
}

func (s *CatalogService) DeleteCatalog(ctx context.Context, req *pb.DeleteCatalogItemRequest) (*pb.CatalogItem, error) {
	// Implement your method here
	return &pb.CatalogItem{}, nil
}
