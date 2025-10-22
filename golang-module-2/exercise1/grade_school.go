package exercise1

import (
	"fmt"
	"sort"
)

type School struct {
	studentList map[int][]string
}

func Grade_school() {
	school := School{
		studentList: make(map[int][]string),
	}
	school.AddStudent("Avinash", 2)
	school.AddStudent("Vinay", 1)
	school.AddStudent("Paladugula", 1)
	school.AddStudent("Marumamula", 3)
	school.AddStudent("Rasagnya", 7)
	school.AddStudent("Vanga", 2)
	school.AddStudent("Akhila", 2)
	school.AddStudent("Guda", 2)

	fmt.Println("Printing the list of student present in grade 2:")
	school.StudentsListInGrade(2)

	fmt.Println("Printing the list of all student sorted by the grade first then sorted by the names:")
	school.GetAllStudents()
}

func (s *School) AddStudent(name string, grade int) {
	s.studentList[grade] = append(s.studentList[grade], name)
	// sort.Strings(s.studentList[grade])
}

func (s *School) StudentsListInGrade(grade int) {
	sort.Strings(s.studentList[grade])
	for _, student := range s.studentList[grade] {
		fmt.Print(student, " ")
	}
	fmt.Println()
}

func (s *School) GetAllStudents() {
	var grades []int
	for g := range s.studentList {
		grades = append(grades, g)
	}
	sort.Ints(grades)
	for _, grade := range grades {
		s.StudentsListInGrade(grade)
	}
}
