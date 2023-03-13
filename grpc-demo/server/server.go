package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"io"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("client.SayHello resp: %s", r.Name)
	return &pb.HelloReply{Message: "hello.world"}, nil
}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	log.Printf("client.SayHello resp: %s", r.Name)
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}
	return nil
}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "say.record"})
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}
	return nil
}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	log.Printf("SayRoute")
	n := 0
	for {
		err := stream.Send(&pb.HelloReply{Message: "hello.list"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("resp: %v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)

}
