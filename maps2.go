package main

import "fmt"

func main()  {
	
	studentAge := make(map{string} int)

	studentAge["rahman"] = 25
	studentAge["khan"] = 20
	studentAge["ram"] = 22
	studentAge["siva"] = 24

	fmt.Println(studentAge["khan"])

}

