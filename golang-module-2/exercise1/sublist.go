package exercise1

import "fmt"

func Sublist_exercise(){
	b := []int{1,2,3,4,5}
	a := []int{3,4,5,6}

	answer := sublist_check(a,b)

	switch answer{
	case "Equal":
		fmt.Println("Both are same")
	case "Sub-list":
		fmt.Println("A is sublist of B")
	case "Super-list":
			fmt.Println("A is superlist of B")
	default:
		fmt.Println("Both are not equal")
	}
}

func sublist_check(a, b []int) string{
	if isEqual(a,b){
		return "Equal"
	}else if isSublist(a,b){
		return "Sub-list"
	}else if isSublist(b,a){
		return "Super-list"
	}
	return "Not Equal"
}

func isEqual(a, b []int) bool{
	if len(a)!=len(b){
		return false
	}
	for i:=0 ; i<len(a) ; i++{
		if a[i]!=b[i]{
			return false
		}
	}
	return true;
}
func isSublist(a,b[]int) bool{
	if len(a)==0 {
		return true
	}
	if len(a)>len(b){
		return false
	}
	for i:=0;i<=len(b)-len(a);i++{
		not_same:=false
		for j:=0 ; j<len(a);j++{
			if a[j]!=b[i+j]{
				not_same=true
				break
			}
		}
		if !not_same{
			return true
		}
	}
	return false
}