package main

import "fmt"

func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	HeapPermutation(a, len(a))
}
