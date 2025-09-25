package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "grpc-employee/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	for {
		fmt.Println("Choose your option:")
		fmt.Println("1. Get employee id")
		fmt.Println("2. get employee name and id greater than the value you entered")
		fmt.Println("3. Add 1 or more employee")
		fmt.Println("4. Get specific employee details")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Print("Enter Employee ID: ")
			var id int32
			fmt.Scanln(&id)
			resp, _ := client.GetEmployee(context.Background(), &pb.EmployeeRequest{Id: id})
			fmt.Printf("You asked for: ID=%d, Name=%s\n", resp.Id, resp.Name)
		case 2:
			fmt.Println("If you enter particular Employee ID I'll give you all employees with greater or equal to that ID")
			fmt.Print("Enter minimum Employee ID: ")
			var id int32
			fmt.Scanln(&id)
			stream, _ := client.GetEmployees(context.Background(), &pb.EmployeeRequest{Id: id})
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				}
				fmt.Printf("ID=%d, Name=%s\n", resp.Id, resp.Name)
			}
		case 3:
			stream, _ := client.AddEmployees(context.Background())
			fmt.Println("Enter employees (id name), enter '0 exit' to finish:")
			for {
				var id int32
				var name string
				fmt.Scan(&id)
				fmt.Scan(&name)
				stream.Send(&pb.Employee{Id: id, Name: name})
				if id == 0 && name == "exit" {
					break
				}
			}
			resp, _ := stream.CloseAndRecv()
			fmt.Printf("Added %d employees\n", resp.Count)
		case 4:
			stream, _ := client.BidirectionalGetEmployees(context.Background())
			go func() {
				for {
					resp, err := stream.Recv()
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Server: ID=%d, Name=%s\n", resp.Id, resp.Name)
				}
			}()
			fmt.Println("Enter Employee IDs (0 to exit):")
			for {
				var id int32
				fmt.Scanln(&id)
				if id == 0 {
					stream.CloseSend()
					time.Sleep(time.Second)
					break
				}
				stream.Send(&pb.EmployeeRequest{Id: id})
			}
		default:
			return
		}
	}
}
