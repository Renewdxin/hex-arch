package grpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"hex-arch/internal/adapters/app/api"
	"hex-arch/internal/adapters/core/arithmetic"
	"hex-arch/internal/adapters/framework/left/grpc/proto/pb"
	"hex-arch/internal/adapters/framework/right/db"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	dbaseAdapter, err := db.NewAdapter("mysql", "admin:password")
	if err != nil {
		fmt.Println(err)
	}
	core := arithmetic.NewAdapter()
	applicationAPI := api.NewAdapter(dbaseAdapter, core)
	grpcAdapter := NewAdapter(applicationAPI)

	pb.RegisterArithmeticServiceServer(grpcServer, *grpcAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error : %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGrpcConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{
		A: 1,
		B: 2,
	}
	answer, err := client.Add(ctx, req)
	if err != nil {
		t.Fatalf("failed to get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(2))

}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{
		A: 1,
		B: 2,
	}
	answer, err := client.Subtract(ctx, req)
	if err != nil {
		t.Fatalf("failed to get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(1))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{
		A: 1,
		B: 2,
	}
	answer, err := client.Multiply(ctx, req)
	if err != nil {
		t.Fatalf("failed to get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(2))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{
		A: 1,
		B: 2,
	}
	answer, err := client.Divide(ctx, req)
	if err != nil {
		t.Fatalf("failed to get addition: %v", err)
	}
	require.Equal(t, answer.Value, int32(0))
}
