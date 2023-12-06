package grpc

import (
	"context"
	"google.golang.org/grpc"
	"hex-arch/internal/adapters/framework/left/grpc/proto/pb"
	"hex-arch/internal/ports"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func (grpca Adapter) Add(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	value, _ := grpca.api.GetAddition(parameters.A, parameters.B)
	return &pb.Answer{
		Value: value,
	}, nil
}

func (grpca Adapter) Subtract(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	value, _ := grpca.api.GetSubtraction(parameters.A, parameters.B)
	return &pb.Answer{
		Value: value,
	}, nil
}

func (grpca Adapter) Multiply(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	value, _ := grpca.api.GetMultiplication(parameters.A, parameters.B)
	return &pb.Answer{
		Value: value,
	}, nil
}

func (grpca Adapter) Divide(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	value, _ := grpca.api.GetDivision(parameters.A, parameters.B)
	return &pb.Answer{
		Value: value,
	}, nil
}

func (grpca Adapter) mustEmbedUnimplementedArithmeticServiceServer() {

}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	arithmeticService := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticService)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
