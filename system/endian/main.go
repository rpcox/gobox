// determine endian
package main

import (
    "fmt"
    "golang.org/x/sys/cpu"
)

func main() {
    if cpu.IsBigEndian {
        fmt.Println("big")
    } else {
        fmt.Println("little")
    }
}

