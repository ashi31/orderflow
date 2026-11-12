package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ashi31/orderflow/gen/orderflowpb"
)

type server struct {
	pb.UnimplementedOrderEventServiceServer
}

func (s *server) SubmitEvent(ctx context.Context, req *pb.SubmitEventRequest) (*pb.SubmitEventResponse, error) {
	e := req.GetEvent()

	log.Printf("Line 1 received event: id=%s order=%s type=%s amount=%d",
		e.GetEventId(), e.GetOrderId(), e.GetEventType(), e.GetAmountMinor())

	return &pb.SubmitEventResponse{EventId: e.GetEventId(), Accepted: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderEventServiceServer(grpcServer, &server{})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}