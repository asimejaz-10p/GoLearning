package main

import "fmt"

func updateName(p *string) {
	*p = "New " + *p
}

type student struct {
	name string
	age  int
}

func createStudentRecord(name string, age int) student {
	s := student{
		name,
		age,
	}
	return s
}

func (s *student) formatData() string {
	var formattedData string
	formattedData += fmt.Sprintf("Student: %v\nAge: %v", s.name, s.age)
	return formattedData
}
func main() {

	student := createStudentRecord("Asim Ejaz", 26)
	fmt.Println(student.formatData())

}
