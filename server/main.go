package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/tanvir/world-clock/proto"
	"google.golang.org/grpc"
)

type timeServer struct {
	pb.UnimplementedTimeServiceServer
}

func (s *timeServer) GetTime(ctx context.Context, req *pb.TimeZoneRequest) (*pb.TimeResponse, error) {
	loc, err := time.LoadLocation(req.Name)
	if err != nil {
		return nil, err
	}

	now := time.Now().In(loc)
	formattedTime := now.Format("2006-01-02 15:04:05")

	resp := &pb.TimeResponse{
		Time: formattedTime,
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &timeServer{})

	log.Println("Serving time requests on port 8081")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
