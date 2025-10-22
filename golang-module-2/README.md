Golang Module 2 Activities

Basic File Reader exercise :
Steps followed:
- Take the path of the file from the console into the string variable "filePath" using the fmt.Scanln(&filePath)
- To open the file, use the method os.Open(filePath) which returns the file text and an error.
- If there is some problem opening the file, it returns the error. So if error occurs, I printed the error and stopped the execution of the program.
- Since the file needs to be closed at any cost I used the defer keyword and call the Close() method
- To scan through the file and print the content use the bufio.NewScanner() method and let that be stored in scanner variable.
- To scan the file line by line and print use the scanner.Scan() and scanner.Text() methods.
scanner.Scan() method checks if there is still some content from the current pointer
scanner.Text() gives the text of the current line and moves the pointer to the next line.
- If there occurs some error  in between while reading the file we can get it using scanner.Err() method.
==> To run the code locally:
- since the code is in exercise2 package, import the package to the main package and call the File_reader() function from the main function
- Run the program using "go run main.go"
- It'll ask for the path of the file. (I entered : /home/beautifulcode/Documents/newDocuments/TIL's/Sep)
- When the path is correct it prints the information from the file. 
- When I gave wrong path it displays the error.
(Handling errors also part of the execution explaination)