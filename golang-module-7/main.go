package main

import "fmt"

type Queue[T any] struct {
	values []T
}

func main() {
	var option int
	q := Queue[any]{}
ForLoop:
	for {
		switch option {
		case 0:
			printOptions()
		case 1:
			fmt.Println("Enter value to be pushed: ")
			var input string
			fmt.Scanln(&input)
			var val any=input
			q.enqueue(val)
		case 2:
			val, exists := q.dequeue()
			if !exists {
				fmt.Println("Queue is empty")
			} else {
				fmt.Println(val)
			}
		case 3:
			peekValue, flag := q.peek()
			if !flag{
				fmt.Println("Queue is empty")
			}else{
				fmt.Println("Peek value is: ",peekValue)
			}
		default:
			break ForLoop
		}
		fmt.Print("Enter the option: ")
		fmt.Scanln(&option)
	}

}
func printOptions() {
	fmt.Println("0 : Print options")
	fmt.Println("1 : Enqueue")
	fmt.Println("2 : Dequeue")
	fmt.Println("Default : End the operations")
}
func (q *Queue[T]) enqueue(value T) {
	q.values = append(q.values, value)
}
func (q *Queue[T]) dequeue() (T, bool) {
	if len(q.values) == 0 {
		var val T
		return val, false
	}
	pop := q.values[0]
	q.values = q.values[1:]
	return pop, true
}
func (q *Queue[T]) peek() (T,bool){
	var val T
	if len(q.values)==0{
		return val,false
	}
	val = q.values[0]
	return val,true
}