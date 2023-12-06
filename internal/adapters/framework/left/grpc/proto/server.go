package proto

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
	//TODO implement me
	panic("implement me")
}

func (grpca Adapter) Subtract(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	//TODO implement me
	panic("implement me")
}

func (grpca Adapter) Multiply(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	//TODO implement me
	panic("implement me")
}

func (grpca Adapter) Divide(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	//TODO implement me
	panic("implement me")
}

func (grpca Adapter) mustEmbedUnimplementedArithmeticServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error
	_, err = net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	arithmeticService := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, &arithmeticService)

}
