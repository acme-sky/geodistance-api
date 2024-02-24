package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"slices"

	pb "github.com/acme-sky/geodistance-api/pkg/distance/proto"
	"google.golang.org/grpc"
	"googlemaps.github.io/maps"
)

var (
	port    = flag.Int("port", 50051, "Server port")
	mclient *maps.Client
)

// `server` implements `distance.DistanceServer`
type server struct {
	pb.UnimplementedDistanceServer
}

// In a lazy way initializes the `mclient` global variable, used to manage a
// Google Maps Client from a `MAPS_KEY` env var
func GetMapsClient() *maps.Client {
	mapsAPIKey := os.Getenv("MAPS_KEY")
	if len(mapsAPIKey) == 0 {
		log.Fatalf("You must set `MAPS_KEY` variable")
	}

	client, err := maps.NewClient(maps.WithAPIKey(mapsAPIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return client
}

// `FindDistance` implements `distance.DistanceServer`
func (s *server) FindDistance(ctx context.Context, in *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	var origin pb.MapPosition = *in.GetOrigin()
	var destination pb.MapPosition = *in.GetDestination()
	log.Printf("Search distance between (%v, %v) and (%v, %v)", origin.GetLatitude(), origin.GetLongitude(), destination.GetLatitude(), destination.GetLongitude())

	if mclient == nil {
		mclient = GetMapsClient()
	}
	matrix := &maps.DistanceMatrixRequest{
		Origins:      []string{fmt.Sprintf("%f,%f", origin.GetLatitude(), origin.GetLongitude())},
		Destinations: []string{fmt.Sprintf("%f,%f", destination.GetLatitude(), destination.GetLongitude())},
		Units:        maps.UnitsMetric,
	}
	resp, err := mclient.DistanceMatrix(context.Background(), matrix)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	var distances []int32

	for _, row := range resp.Rows {
		distances = append(distances, int32(row.Elements[0].Distance.Meters))
	}

	return &pb.DistanceResponse{Origin: resp.OriginAddresses[0], Destination: resp.DestinationAddresses[0], Distance: slices.Min(distances)}, nil
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
