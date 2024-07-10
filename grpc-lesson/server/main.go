package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListeFiles was invoked")

	dir := "/workspaces/grpc/grpc-lesson/storage"

	paths, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, path := range paths {
		if !path.IsDir() {
			fileNames = append(fileNames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: fileNames,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("Server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
