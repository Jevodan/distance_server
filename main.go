package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	api "github.com/Jevodan/proto/distance"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "port number to connect")
)

type server struct {
	api.UnimplementedDistanceServer
}

func (s *server) GetDistance(ctx context.Context, req *api.Points) (*api.Dist, error) {
	distance := math.Sqrt(math.Pow(req.B.X-req.A.X, 2) + math.Pow(req.B.Y-req.A.Y, 2))
	return &api.Dist{Result: distance}, nil
}

func main() {
	flag.Parse()
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	api.RegisterDistanceServer(s, &server{})
	log.Printf("server listening at %v", conn.Addr())
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
