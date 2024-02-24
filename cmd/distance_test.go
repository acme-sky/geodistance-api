package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/acme-sky/geodistance-api/pkg/distance/proto"
)

var lis *bufconn.Listener

func serv(ctx context.Context) (pb.DistanceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterDistanceServer(baseServer, &server{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewDistanceClient(conn)

	return client, closer
}

func TestFindDistance(t *testing.T) {
	ctx := context.Background()
	client, closer := serv(ctx)
	defer closer()

	type expectation struct {
		out *pb.DistanceResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.DistanceRequest
		expected expectation
	}{
		"Success": {
			in: &pb.DistanceRequest{
				Position1: &pb.MapPosition{Latitude: 44.4970717, Longitude: 11.3535314},
				Position2: &pb.MapPosition{Latitude: 37.5257372, Longitude: 15.0702872}},
			expected: expectation{
				out: &pb.DistanceResponse{
					Distance: 42.1,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.FindDistance(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Distance != out.Distance {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}
