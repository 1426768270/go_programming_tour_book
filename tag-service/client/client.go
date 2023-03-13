package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "tag-service/proto"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "java"})
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string,
	opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
