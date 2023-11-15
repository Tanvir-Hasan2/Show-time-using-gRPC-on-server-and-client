package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tanvir/world-clock/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)

	req := &pb.TimeZoneRequest{
		Name: "Europe/Dublin",
	}
	resp, err := client.GetTime(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to get time: %v", err)
	}

	fmt.Println("Current time in Europe/Dublin:", resp.Time)
}
