package main

import "fmt"

// ReverseBits reverses the bits in a 64-bit unsigned integer
func ReverseBits(x uint64) uint64 {
//loop over every bit of x, extract only bit at position i, all other bits being 0.
//shift it( n-i) and add it to result through OR operator

    result := uint64(0)
    for i := 0; i < 64; i++ {
        bit := (x >> i) & 1
        result |= bit << (63 - i)
    }
    return result
}

func main() {
    x := uint64(0b111000)
    fmt.Printf("x is %064b\n", x)
    x = ReverseBits(x)
    fmt.Printf("x becomes %064b\n", x)
}
