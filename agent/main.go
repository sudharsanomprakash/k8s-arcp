package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"github.com/sudharsanomprakash/k8s-arcp/internal"
	pb "github.com/sudharsanomprakash/k8s-arcp/proto"
)

type server struct {
	pb.UnimplementedARCPServer
}

func (s *server) Exchange(ctx context.Context, in *pb.ResourceStatus) (*pb.ResourceStatus, error) {
	log.Printf("Received status from %s: CPU %.2f%%, Mem %dMiB", in.PodID, in.CpuUsage, in.MemUsage)
	
	// Basic suggestion logic
	suggested := "normal"
	if in.CpuUsage > 80 {
		suggested = "scale_up"
	} else if in.CpuUsage < 20 {
		suggested = "scale_down"
	}
	
	metrics, _ := internal.Collect()
	return &pb.ResourceStatus{
		PodID:          os.Getenv("POD_ID"),
		Namespace:      os.Getenv("POD_NAMESPACE"),
		CpuUsage:       float32(metrics.CPUUsage),
		MemUsage:       metrics.MemUsage,
		Timestamp:      uint64(metrics.Timestamp),
		RequestedCPU:   50.0,
		RequestedMem:   1024,
		SuggestedAction: suggested,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterARCPServer(s, &server{})
	log.Println("ARCP agent is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

