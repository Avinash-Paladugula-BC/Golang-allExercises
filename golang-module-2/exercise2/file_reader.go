package file_reader

import (
	"bufio"
	"fmt"
	"os"
)
// /home/beautifulcode/Documents/newDocuments/TILs/Sep
func File_reader() {
	defer func(){
		if r := recover(); r!=nil{
			fmt.Printf("Error occured: %v\n",r)
		}
	}()
	fmt.Println("Enter the file path: ")
	var filePath string
	fmt.Scanln(&filePath)

	fileText, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf("Error while opening file ",err)
	}
	defer fileText.Close()
	scanner := bufio.NewScanner(fileText)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
