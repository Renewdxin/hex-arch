package main

import (
	"fmt"
	"hex-arch/internal/adapters/app/api"
	"hex-arch/internal/adapters/core/arithmetic"
	"hex-arch/internal/adapters/framework/left/grpc"
	"hex-arch/internal/adapters/framework/right/db"
	"hex-arch/internal/ports"
)

func main() {
	var core ports.ArithmeticPort
	var err error
	var dbaseAdapter ports.DBPort
	var arithAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbaseAdapter, err = db.NewAdapter("mysql", "admin:password")
	arithAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(arithAdapter, dbaseAdapter)
	grpcAdapter = grpc.NewAdapter(appAdapter)
	grpcAdapter.Run()

	core = arithmetic.NewAdapter()
	result, err := core.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
