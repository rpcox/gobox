package main

import "testing"


func Reset(a, reset *[]byte) {
        *a = *reset
}

func BenchmarkReset(b *testing.B) {
	var reset = make([]byte, 1024)
	var a = make([]byte, 1024)

	for b.Loop()  {
        	a[10] = 10
        	a[20] = 20
		Reset(&a, &reset)
	}
}

