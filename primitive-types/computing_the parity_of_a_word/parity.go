package main

import (
    "fmt"
)

func xorParitySingle(x uint64) uint64 {
    fmt.Printf("%064b\n", x)
    x ^= x >> 32
    fmt.Printf("%064b\n", x)
    x ^= x >> 16
    fmt.Printf("%064b\n", x)
    x ^= x >> 8
    fmt.Printf("%064b\n", x)
    x ^= x >> 4
    fmt.Printf("%064b\n", x)

    x ^= x >> 2
    fmt.Printf("%064b\n", x)

    x ^= x >> 1
    fmt.Printf("%064b\n", x)
    return x & 1
}

func main() {
    x := uint64(0xFFFF) // Three 1s → parity should be odd → expect 1
    fmt.Println(x)
    fmt.Println(xorParitySingle(x)) // Output: 1
}

