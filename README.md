# ACME Sky - GeoDistance service

This repo refers to the GeoDistance service used by ACME Sky to obtain distance
between two locations.
GeoDistance uses the Google Maps Distance Matrix API as a backend to get distance 
info in meters.

## Maps API

The API key can be restricted to "Distance Matrix API", "Directions API" and
"Gecoding API".

## gRPC

gRPC is used as the protocol for communicating with this API. Here's a small
client:

```go
// ...
c := pb.NewDistanceClient(conn)

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
r, err := c.FindDistance(ctx, &pb.DistanceRequest{
    Origin: &pb.MapPosition{Latitude: 44.4969, Longitude: 11.3564347},
    Destination: &pb.MapPosition{Latitude: 37.5257372, Longitude: 15.0702872}})
if err != nil {
    log.Fatalf("Could not find distance: %v", err)
}
log.Printf("%d", r.GetDistance())
```

Also, you can find the map position from an address:

```go
// ...
c := pb.NewDistanceClient(conn)

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
r, err := c.FindGeometry(ctx, &pb.AddressRequest{Address: "Viale Andrea Doria 6, Catania"})
if err != nil {
    log.Fatalf("Could not find distance: %v", err)
}
log.Printf("%v", r)
```

## Setup

You need to set up `MAPS_KEY` env variable for `go run cmd/distance.go`.
If you change something in `*.proto` files, please run `make proto`.

## Build

You can use Docker for a fresh build.

```
docker build -t acmesky-geodistance-api .

docker run -e MAPS_KEY=.... --name geodistance -p 50051:50051 --network acmesky acmesky-geodistance-api:latest
```
