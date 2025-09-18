package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func main() {

	// get the id's first and store them in ids[]
	// user_ids := make([]int ,0)
	user_ids, err := getAllUsers()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	// create channel of size 5
	ch := make(chan int, 5)
	// call them
	var wg sync.WaitGroup
	for idx, user_id := range user_ids {
		fmt.Println(idx)
		ch <- user_id
		wg.Add(1)
		go func(user_id int) {
			user, err := getUserByID(user_id)
			if err != nil {
				fmt.Println("error: ", err)
			}
			fmt.Println(user)
			<-ch
			wg.Done()
		}(user_id)
	}
	wg.Wait()

}

func getAllUsers() ([]int, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}
	user_ids := make([]int, 0)
	for _, user := range users {
		user_ids = append(user_ids, user.ID)
	}
	return user_ids, nil
}

func getUserByID(id int) (User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}