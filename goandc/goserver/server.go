package main

//go:generate protoc -I../proto/query --go_out=plugins=grpc:../proto/query query.proto
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"strconv"

	pb "github.com/hu-univ-golang/golang-elte-2020-public/goandc/proto/query"
)

type service struct{}

func (s *service) Query(_ context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	l, err := strconv.Atoi(req.Query)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%q had problems: %s", req.Query, err)
	}
	resp := &pb.QueryResponse{}
	for i := 0; i < l; i++ {
		resp.Response = append(resp.Response, fmt.Sprintf("r-%d", i))
	}
	return resp, nil
}

const port = ":54321"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQueryServiceServer(s, &service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
