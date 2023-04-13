package main

import "fmt"

func main() {
	e := Employee{
		name:     "Bùi Hiên",
		salary:   3000,
		currency: "$",
	}

	e.displaySalary()

	fmt.Println("employee name before changeName : ", e.name)
	e.changeName("New Name")
	fmt.Println("employee name after changeName : ", e.name)
	fmt.Println("employee salary before changeSalary : ", e.salary)
	e.changeSalary(4000)
	fmt.Println("employee salary after changeSalary : ", e.salary)
}

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %d%s", e.name, e.salary, e.currency)
}

func (e Employee) changeName(name string) {
	e.name = name
}

func (e *Employee) changeSalary(salary int) {
	e.salary = salary
}
