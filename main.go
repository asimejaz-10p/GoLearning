// package main

// import (
// 	_ "GoLearning/services"
// 	"fmt"
// )

// func updateName(p *string) {
// 	*p = "New " + *p
// }

// type student struct {
// 	name string
// 	age  int
// }

// func createStudentRecord(name string, age int) student {
// 	s := student{
// 		name,
// 		age,
// 	}
// 	return s
// }

// func (s *student) formatData() string {
// 	var formattedData string
// 	formattedData += fmt.Sprintf("Student: %v\nAge: %v", s.name, s.age)
// 	return formattedData
// }
// func main() {

// 	//student := createStudentRecord("Asim Ejaz", 26)
// 	//fmt.Println(student.formatData())

// 	name := "Asim"

// 	switch {
// 	case name == "Asim":
// 		fmt.Println("I am Asim 123")
// 	case name == "Asim":
// 		fmt.Println("I am Asim again")
// 	case name == "Asad":
// 		fmt.Println("I am Asad")

// 	}
// }

package main

import (
	Services "GoLearning/services"
	"fmt"
)

func main() {
	// encodedValue := "https://app.boostlingo.com/api/web/call/call-log?q=%7B%22callStatusId%22%3Anull%2C%22communicationType%22%3A0%2C%22companyAccountId%22%3A41082%2C%22consumerId%22%3Anull%2C%22dateSince%22%3A%222021-11-10T19%3A40%3A35.77863324Z%22%2C%22dateTill%22%3A%222021-11-12T19%3A40%3A35.77863324Z%22%2C%22isInternal%22%3Afalse%2C%22pageIndex%22%3A1%2C%22pageSize%22%3A25%2C%22requestorId%22%3A0%2C%22searchString%22%3Anull%2C%22sortBy%22%3A%22timeConnected%22%2C%22sortDirection%22%3A2%2C%22userAccountId%22%3A0%7D"
	// decodedValue, err := url.QueryUnescape(encodedValue)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(decodedValue)
	numbers := []int{3, 6, 1, 8}

	fmt.Printf("Sum is %v\n", addNumbers(numbers))
	fmt.Println(Services.DisplayData())

}
