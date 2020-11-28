package main

import "fmt"

func main() {
	// Make new maps.
	fmt.Println(make(map[string]int)) // map[]
	m := make(map[string]int, 5)
	fmt.Println(m, len(m)) // map[] 1
	m["apple"] = 2000
	m["nokia"] = 1995
	fmt.Println(m, len(m)) // map[apple:2000 nokia:1995] 2

	// Make new slices.
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [1 1 1] 3 5
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [1 1] 2 2
}
