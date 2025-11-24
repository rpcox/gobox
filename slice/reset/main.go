package main

import "fmt"

func main() {
	b := make([]byte, 32)
	a := make([]byte, 32)
	fmt.Printf("a @ p=%p %v\n", &a, a) // [0 0 .... 0]
	fmt.Printf("b @ p=%p %v\n", &b, b) // [0 0 .... 0]

	a[10] = 10
	a[20] = 20
	fmt.Println(a) // [0 ... 10 ... 20 .... 0]

	// Reset array
	a = b
	fmt.Printf("a @ p=%p %v\n", &a, a) // [0 0 .... 0]
}
