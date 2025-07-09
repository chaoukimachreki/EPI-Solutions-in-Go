package main

import (
    "fmt"
)

// SwapBits swaps bits at positions i and j in x if they differ.
// i and j must be between 0 and 63 inclusive.
func SwapBits(x uint64, i, j int) uint64 {
    if i < 0 || i >= 64 || j < 0 || j >= 64 {
        panic("bit index out of range")
    }

    if ((x >> i) & 1) != ((x >> j) & 1) {
        bitMask := (uint64(1) << i) | (uint64(1) << j)
	fmt.Printf("bit mask is : %064b\n",bitMask)
        x ^= bitMask
    }
    return x
}

func main() {
    x := uint64(0xABBFECCEFBABCE0C)
    fmt.Printf("x in binary :%064b\n",x)
    result := SwapBits(x, 2, 1)
    fmt.Printf("after swap x becomes: %064b\n", result)
}

