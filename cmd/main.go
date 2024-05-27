
package main

import (
	"context"
	"flag"

	// "github.com/gusmann/go-bzzt/gateway"
	"github.com/gusmann/go-bzzt/server"
	"google.golang.org/grpc/grpclog"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	openAPIDir = flag.String("openapi_dir", "examples/internal/proto/examplepb", "path to the directory which contains OpenAPI definitions")
)

func main() {
	flag.Parse()

	// ctx := context.Background()
	// opts := gateway.Options{
	// 	Addr: ":8080",
	// 	GRPCServer: gateway.Endpoint{
	// 		Network: *network,
	// 		Addr:    *endpoint,
	// 	},
	// 	OpenAPIDir: *openAPIDir,
	// }
	// if err := gateway.Run(ctx, opts); err != nil {
	// 	grpclog.Fatal(err)
	// }
	if err := server.Run(context.Background()); err != nil {
		grpclog.Fatal(err)
	}
}