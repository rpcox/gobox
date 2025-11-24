package main

import "testing"


func Reset(a *[]byte) {
        for i := 0; i < len(*a); i++ {
		(*a)[i] = 0
	}
}

func BenchmarkReset(b *testing.B) {
	var a = make([]byte, 1024)

	for b.Loop()  {
        	a[10] = 10
        	a[20] = 20
		Reset(&a)
	}
}

