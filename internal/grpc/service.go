package grpc

import (
	"context"

	"github.com/lubie-koty/rpc-compute-service-complex/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-complex/protos"
)

type GRPCService struct {
	pb.UnimplementedComplexComputeServer
	service types.MathService
}

func NewGRPCService(service types.MathService) *GRPCService {
	return &GRPCService{
		service: service,
	}
}

func (s *GRPCService) Sqrt(ctx context.Context, req *pb.UnaryRequest) (*pb.OperationResponse, error) {
	result := s.service.Sqrt(req.GetNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) Abs(ctx context.Context, req *pb.UnaryRequest) (*pb.OperationResponse, error) {
	result := s.service.Abs(req.GetNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) Power(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	result := s.service.Power(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) Log(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	result := s.service.Log(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) Round(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	result := s.service.Round(req.GetFirstNumber(), int64(req.GetSecondNumber()))
	return &pb.OperationResponse{Result: result}, nil
}
