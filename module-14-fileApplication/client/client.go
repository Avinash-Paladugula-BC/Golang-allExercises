package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "file-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func uploadFiles(client pb.FileServiceClient, files []string) {
	for _, file := range files {
		text, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("We can't find the file provided path %s: %v", file, err)
		}
		req := pb.FileData{
			Filename: file,
			Content:  text,
		}

		resp, err := client.UploadFile(context.Background(), &req)
		if err != nil {
			log.Fatalf("Sorry to say upload failed: %v", err)
		}
		fmt.Println("Upload:", resp.Message)
	}
}

func downloadFiles(client pb.FileServiceClient, files []string) {
	for _, file := range files {
		req := &pb.DownloadRequest{Filename: file}
		res, err := client.DownloadFile(context.Background(), req)
		if err != nil {
			log.Fatalf("Download process failed brother! Issue : %v", err)
		}
		err = os.WriteFile("client_"+res.Filename, res.Content, 0644)
		if err != nil {
			log.Fatalf("failed to save downloaded file: %v", err)
		}
		fmt.Println("Hip hup hurray!! Downloaded the file:", res.Filename)
	}
}

func getMetadata(client pb.FileServiceClient, filename string) {
	req := &pb.MetadataRequest{Filename: filename}
	resp, err := client.GetFileMetadata(context.Background(), req)
	if err != nil {
		log.Fatalf("metadata error: %v", err)
	}
	for _, f := range resp.Files {
		fmt.Printf("Filename: %s \nSize: %d \nUploaded: %s\n", f.Filename, f.Size, f.UploadTime)
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("The way you need to run the client file is: ")
		fmt.Println("go run client/client.go [upload/download/metadata] space_separated_file_names")
		return
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFileServiceClient(conn)

	option := os.Args[1]
	files := os.Args[2:]

	switch option {
	case "upload":
		uploadFiles(client, files)
	case "download":
		downloadFiles(client, files)
	case "metadata":
		// If No file name is included, it means we need to print log details of all files
		filename := ""
		if len(files) != 0 {
			filename = files[0]
		}
		getMetadata(client, filename)
	default:
		fmt.Println(option," cannot be performed, sorry :)")
	}
}
