MODULE 4 EXERCISE - ERROR HANDLING AND TYPE CONVERSION

- Since the input needs to be taken from the command line and if the input values are not passed through the command line, then we need to handle it as an error. So first I've written the recover function using the defer keyword.
- To take the input we use string and the flag name is numbers
- The variale needs to be parsed using the flag.Parse() method
- Since the input is in string format which is numbers separated by commas, I've written a funciton that returns the list of strings considering the comma as delimeter
- Now to store the sum of all the values I've created a integer value named sum
- By iterating through the list of strings we try to convert it to integer.
- While converting the string into the number if error occurs then I've created a panic, else it will be add to the sum vaiable
- Finally print the sum variable
====Running the code locally:
- Instead of using the command : "go run main.go" I used : "go run main.go -numbers=2,1,4(comma separated numbers)"
- Here the numbers indicate the flag name followed by the value.
====Type conversion and error handling
- To convert the number which is in the form of string use "strconv.Atoi(str_variable)"
- To handle the error any point I used the recover() and printed the error. When ever a string is given in between the numbers or no numbers are passed in the command line, the program creates a panic.