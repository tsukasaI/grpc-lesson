package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"net"
	"os"
	"time"

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

func (*server) Download(req *pb.DownloadRequest, stream pb.FileService_DownloadServer) error {
	fmt.Println("Download was invoked")

	filename := req.GetFilename()

	path := fmt.Sprintf("/workspaces/grpc/grpc-lesson/storage/%s", filename)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}
		time.Sleep(1 * time.Second)
	}
	return nil
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
