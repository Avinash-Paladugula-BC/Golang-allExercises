package file_reader

import (
	"bufio"
	"fmt"
	"os"
)
// /home/beautifulcode/Documents/newDocuments/TIL's/Sep
func File_reader() {
	fmt.Println("Enter the file path: ")
	var filePath string
	fmt.Scanln(&filePath)

	fileText, err := os.Open(filePath)
	if err != nil {
		fmt.Println("May be the entered file path is wrong. Error: ", err)
		return
	}
	defer fileText.Close()
	scanner := bufio.NewScanner(fileText)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}


}
