package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	// "math/rand"
)

func main() {
	http.HandleFunc("/process", processHandler)
	fmt.Println("Server will be running on localhost:8080/process , please open it in browser")


	// this will start the server and receives the request
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// to stop after certain time we use WithTimeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	//this is the time taken to complete the task--simulation
	// randomInt := 6 * time.Second 
	randomInt := 3 * time.Second
	// randomInt,_ := rand.Int(rand.Reader, big.NewInt(10))
	// taskDuration := randomInt * time.Second
	// fmt.Printf("The server will be timed out if the %d >= 5",randomInt)
	
	select {
	case <-time.After(randomInt):
		// Task finished before timeout or cancellation
		fmt.Println("Task completed successfully!")
		// fmt.Println("Task completed successfully")
		
	case <-ctx.Done():
		// Context ended: either timeout or client cancel
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			fmt.Println("Task cancelled since timeout reached")
			} else if err == context.Canceled {
				http.Error(w, "Request cancelled by client", http.StatusRequestTimeout)
				fmt.Println("Task cancelled since connection was closed")
			}
		}
	}
	