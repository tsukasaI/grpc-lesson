package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connext: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	// callListFiles(client)
	callDownload(client)

}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}

func callDownload(client pb.FileServiceClient) {
	req := &pb.DownloadRequest{Filename: "name.txt"}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Response from download(Bytes): %v", res.GetData())
		log.Printf("Response from download(string): %v", string(res.GetData()))
	}
}
