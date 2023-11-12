package main

import (
	"os"

	reverse "github.com/SaveliiYam/grpc_simple_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	args := os.Args
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := reverse.NewReverseClient(conn)
}
