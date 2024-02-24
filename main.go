package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/acme-sky/geodistance-api/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

// `server` implements `distance.DistanceServer`
type server struct {
	pb.UnimplementedDistanceServer
}

// `FindDistance` implements `distance.DistanceServer`
func (s *server) FindDistance(ctx context.Context, in *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	var pos1 pb.MapPosition = *in.GetPosition1()
	var pos2 pb.MapPosition = *in.GetPosition2()
	log.Printf("Received: (%v, %v) (%v, %v)", pos1.GetLatitude(), pos1.GetLongitude(), pos2.GetLatitude(), pos2.GetLongitude())
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
