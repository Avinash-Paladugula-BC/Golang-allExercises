package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "file-grpc/proto"
	"file-grpc/storage"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (s *server) UploadFile(ctx context.Context, req *pb.FileData) (*pb.UploadStatus, error) {
	err := storage.SaveFile(req.Filename, req.Content)
	if err != nil {
		return &pb.UploadStatus{
			Filename: req.Filename,
			Success:  false,
			Message:  fmt.Sprintf("Unable to save file: %v", err),
		}, nil
	}
	return &pb.UploadStatus{
		Filename: req.Filename,
		Success:  true,
		Message:  "File uploaded successfully",
	}, nil
}

func (s *server) DownloadFile(ctx context.Context, req *pb.DownloadRequest) (*pb.FileData, error) {
	data, err := storage.ReadFile(req.Filename)
	if err != nil {
		return nil, err
	}
	return &pb.FileData{
		Filename: req.Filename,
		Content:  data,
	}, nil
}

func (s *server) GetFileMetadata(ctx context.Context, req *pb.MetadataRequest) (*pb.MetadataResponse, error) {
	metadata, err := storage.GetMetadata(req.Filename)
	if err != nil {
		return nil, err
	}

	resp := &pb.MetadataResponse{}
	for _, md := range metadata {
		resp.Files = append(resp.Files, &pb.FileMetadata{
			Filename:   md.Filename,
			Size:       md.Size,
			UploadTime: md.UploadTime,
		})
	}
	return resp, nil
}

func main() {
	// 50051 port is default port for client server interaction in grpc
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &server{})

	fmt.Println("Server running on port 50051...")
	fmt.Println("before running the client.go ==> ")
	fmt.Println("The way you need to run the client file is: ")
	fmt.Println("go run client/client.go [upload/download/metadata] space_separated_file_names")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
