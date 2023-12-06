package proto

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hex-arch/internal/adapters/framework/left/grpc/proto/pb"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}
	answer, err := grpca.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}
	return ans, nil
}
