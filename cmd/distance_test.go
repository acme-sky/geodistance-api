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
				Origin:      &pb.MapPosition{Latitude: 44.4969, Longitude: 11.356435},
				Destination: &pb.MapPosition{Latitude: 37.525738, Longitude: 15.070287}},
			expected: expectation{
				out: &pb.DistanceResponse{
					Origin:      "Mura Anteo Zamboni, 7, 40126 Bologna BO, Italy",
					Destination: "Viale Andrea Doria, 6, 95125 Catania CT, Italy",
					Distance:    1153265,
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
				if tt.expected.out.Distance != out.Distance && tt.expected.out.Origin != out.Origin || tt.expected.out.Destination != out.Destination {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}
