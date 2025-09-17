package main

import (
	"fmt"
	exercism "module2-exercies/exercise1"
	file "module2-exercies/exercise2"
)

func main() {
	fmt.Println("Exercism Exercises:")

	fmt.Println("\nSublist exercise:")
	exercism.Sublist_exercise()

	fmt.Println("\nWord Coutn exercise:")
	exercism.Word_coun()

	fmt.Println("\nETL Exercise:")
	exercism.Etl()

	fmt.Println("\nGrade School:")
	exercism.Grade_school()

	fmt.Println("===========================================================================================")

	fmt.Println("\n Basic file reader exercise:")
	file.File_reader()

}
