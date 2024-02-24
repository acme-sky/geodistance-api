package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/acme-sky/geodistance-api/distance"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

// `server` implements `distance.DistanceServer`
type server struct {
	pb.UnimplementedDistanceServer
}

// `GetDistance` implements `distance.DistanceServer`
func (s *server) GetDistance(ctx context.Context, in *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	log.Printf("Received: (%v, %v) (%v, %v)", in.GetLat1(), in.GetLon1(), in.GetLat2(), in.GetLon2())
	return &pb.DistanceResponse{Distance: 42.1}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDistanceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
