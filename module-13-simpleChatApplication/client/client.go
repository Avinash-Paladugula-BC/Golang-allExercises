package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	// "time"

	proto "grpc-chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	var username string
	fmt.Scanln(&username)

	for {
		fmt.Println("Choose:")
		fmt.Println("1 - Private Message")
		fmt.Println("2 - Join Chat Room (Server streaming)")
		fmt.Println("3 - Chat Stream (Bidirectional)")
		fmt.Print("Enter choice: ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Print("Enter recipient username: ")
			var to string
			fmt.Scanln(&to)

			fmt.Print("Enter message: ")
			var msg string
			fmt.Scanln(&msg)

			resp, _ := client.SendPrivateMessage(context.Background(), &proto.PrivateMessageRequest{
				From:    username,
				To:      to,
				Message: msg,
			})
			fmt.Println("Server:", resp.Status)

		case "2":
			stream, _ := client.JoinChatRoom(context.Background(), &proto.JoinRequest{Username: username})
			go func() {
				for {
					msg, err := stream.Recv()
					if err != nil {
						log.Println("Stream closed:", err)
						return
					}
					fmt.Printf("%s: %s\n", msg.From, msg.Message)
				}
			}()
			// select {}  --> this will open the server chat until we close
			time.Sleep(60*time.Second) // this will allow for 1 minute

		case "3":
			stream, _ := client.ChatStream(context.Background())
			stream.Send(&proto.ChatMessage{From: username, Message: "joined"})
			go func() {
				for {
					msg, err := stream.Recv()
					if err != nil {
						log.Println("Stream closed:", err)
						return
					}
					fmt.Printf("%s: %s\n", msg.From, msg.Message)
				}
			}()
			for {
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
				stream.Send(&proto.ChatMessage{From: username, Message: text})
			}
		}
	}
}
