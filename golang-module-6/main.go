package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	// http.Get() method will return the pointer to the response
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		panic("Error fetching data")
	}
	defer response.Body.Close()

	// getting the entire content from the page  [ioutil.ReadAll() is depricated so using io.ReadAll()]
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	// the entire information from the body in unmarshalled into result
	var result []User
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	// fmt.Println(result)
	// ...............................................
	// This is for downloading the Avatar from the page by providing the user.ID
	for i, user := range result {
		if i < 100 {
			// Download avatar
			if err := downloadFile(user.Avatar, fmt.Sprintf("user_data/avatar_%d.jpg", user.ID)); err != nil {
				fmt.Printf("Error downloading avatar for user %d: %v\n", user.ID, err)
				continue
			}
		}
	}
	// .....................................................
	dirName := "user_data"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error while creating the directory : ", err)
		fmt.Println("Delete the user_data directory if it is already created")
		return
	}
	fmt.Println("Directory created.......")

	// the below line will convert the slice of struct into the JSON string
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	filePath := filepath.Join(dirName, "users.json")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data: ")
	fmt.Println(string(jsonData))
	fmt.Println("====================================")
	fmt.Println("Slice of User struct format:")
	fmt.Println(result)
	// dirPath := "user_data/users.json"

	// err := os.MkdirAll(dirPath, 0755)
	// if err != nil {
	// 	fmt.Println("Error creating directory:", err)
	// 	return
	// }

	// fmt.Println("Directory created:", dirName)

}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
