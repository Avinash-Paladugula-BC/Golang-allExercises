package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	println("hello world")

	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Unexpected error: ",r)
		}
	}()

	nums_string := flag.String("numbers", "", "Enter the numbers separated by comma")
	flag.Parse()

	if *nums_string == "" {
		fmt.Println("You didnot any numbers while running the main function")
		panic("Numbers not entered")
	}
	nums := splitInput(*nums_string)
	var sum int
	for _, s := range nums {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s is not an integer :(\n", s)
			panic("Not a number")
		}
		sum += number
	}
	fmt.Println("Sum : ",sum)
}

func splitInput(nums_string string) []string{
	var result []string
	start:=0
	for idx,char := range nums_string{
		if char==','{
			result = append(result, nums_string[start:idx])
			start=idx+1
		}
	}
	result = append(result, nums_string[start:])
	return result
}