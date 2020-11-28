package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"rahman":  12000,
		"khan":    15000,
		"suresh":  10000,
		"vignesh": 20000,
		"siva":    25000,
	}
	fmt.Println("Salary of khan is", employeeSalary["khan"])
}
