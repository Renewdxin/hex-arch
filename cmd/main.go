package main

import (
	"fmt"
	"hex-arch/internal/adapters/app/api"
	"hex-arch/internal/adapters/core/arithmetic"
	"hex-arch/internal/adapters/framework/left/grpc"
	"hex-arch/internal/adapters/framework/right/db"
)

func main() {
	dbaseAdapter, err := db.NewAdapter("mysql", "admin:password")
	if err != nil {
		fmt.Println(err)
	}
	defer dbaseAdapter.CloseDBConnection()

	core := arithmetic.NewAdapter()
	applicationAPI := api.NewAdapter(dbaseAdapter, core)
	grpcAdapter := grpc.NewAdapter(applicationAPI)
	grpcAdapter.Run()

	core = arithmetic.NewAdapter()
	result, err := core.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
