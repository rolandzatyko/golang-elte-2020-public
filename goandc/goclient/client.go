package main

//go:generate protoc -I../proto/query --go_out=plugins=grpc:../proto/query query.proto
import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "github.com/hu-univ-golang/golang-elte-2020-public/goandc/proto/query"
)

const port = ":54321"

func main() {
	ctx := context.Background()
	flag.Parse()

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewQueryServiceClient(conn)

	for _, q := range flag.Args() {
		resp, err := client.Query(ctx, &pb.QueryRequest{Query: q})
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}
		fmt.Printf("%q -> %q\n", q, resp.Response)
	}
}
