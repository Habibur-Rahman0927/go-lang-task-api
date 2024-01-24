package structStatement

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func StructUsing() {
	// Pers1 specification
	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

}
