package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "grpc-employee/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
	employees map[int32]string
}

func newServer() *server {
	return &server{employees: make(map[int32]string)}
}

func (s *server) GetEmployee(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	name, ok := s.employees[req.Id]
	if !ok {
		return &pb.EmployeeResponse{Id: req.Id, Name: "Not Found"}, nil
	}
	return &pb.EmployeeResponse{Id: req.Id, Name: name}, nil
}

func (s *server) GetEmployees(req *pb.EmployeeRequest, stream pb.EmployeeService_GetEmployeesServer) error {
	for id, name := range s.employees {
		if id >= req.Id {
			if err := stream.Send(&pb.EmployeeResponse{Id: id, Name: name}); err != nil {
				return err
			}
		}
	}
	return nil
}
func (s *server) AddEmployees(stream pb.EmployeeService_AddEmployeesServer) error {
	count := 0
	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.EmployeeCount{Count: int32(count)})
		}
		if err != nil {
			return err
		}
		if emp.Id == 0 && emp.Name == "exit" {
			return stream.SendAndClose(&pb.EmployeeCount{Count: int32(count)})
		}
		s.employees[emp.Id] = emp.Name
		count++
	}
}

func (s *server) BidirectionalGetEmployees(stream pb.EmployeeService_BidirectionalGetEmployeesServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		name := s.employees[req.Id]
		if name == "" {
			name = "Not available"
		}
		if err := stream.Send(&pb.EmployeeResponse{Id: req.Id, Name: name}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, newServer())

	fmt.Println("Server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
