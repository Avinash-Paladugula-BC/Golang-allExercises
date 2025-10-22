package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	proto "grpc-chat/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedChatServiceServer
	clients map[string]proto.ChatService_JoinChatRoomServer
	streams map[string]proto.ChatService_ChatStreamServer
}

func newServer() *server {
	return &server{
		clients: make(map[string]proto.ChatService_JoinChatRoomServer),
		streams: make(map[string]proto.ChatService_ChatStreamServer),
	}
}

// Unary RPC
func (s *server) SendPrivateMessage(ctx context.Context, req *proto.PrivateMessageRequest) (*proto.PrivateMessageResponse, error) {
	receiver, ok := s.clients[req.To] //if the person to receive is online or not
	if !ok {
		return &proto.PrivateMessageResponse{Status: "User not found"}, nil //--------
	}
	msg := &proto.ChatMessage{From: req.From, Message: req.Message}
	receiver.Send(msg)
	return &proto.PrivateMessageResponse{Status: "Message delivered"}, nil
}

func (s *server) JoinChatRoom(req *proto.JoinRequest, stream proto.ChatService_JoinChatRoomServer) error {
	username := req.Username
	s.clients[username] = stream

	
	time.Sleep(60*time.Second)
	return nil
}

func (s *server) ChatStream(stream proto.ChatService_ChatStreamServer) error {
	firstMsg, err := stream.Recv()
	if err != nil {
		return err
	}
	username := firstMsg.From
	s.streams[username] = stream
	// Sending the message to everyone except the sender
	for _, st := range s.streams { 
		if st != stream {
			st.Send(&proto.ChatMessage{From: "Server", Message: fmt.Sprintf("%s joined", username)})
		}
	}
	for {
		fmt.Println("entered the loop")
		msg, err := stream.Recv()
		fmt.Println(msg)
		if err != nil {
			return err
		}
		// username := msg.From
		if fmt.Sprintf("%v",msg) == "exit" {
			break
		}
		for user, st := range s.streams {
			if user != username {
				st.Send(msg)
			}
		}
		
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterChatServiceServer(grpcServer, newServer())
	log.Println("Server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}