package main

import "fmt"

func main()  {

	Student  :=  make(map{string} int)
    student["rahman"]=25
	fmt.Println(student ["rahman"])
	fmt.Println(len (student))

	student ["khan"]=20
	fmt.Println(len(student))
	delete(student,"khan")
	fmt.Println(len(student))
	
}
