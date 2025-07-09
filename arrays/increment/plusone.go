package main

import (
    "fmt"
)

// PlusOne increments the integer represented by a slice of digits.
func PlusOne(digits []int) []int {
    n := len(digits)
    digits[n-1]++

    for i := n - 1; i > 0 && digits[i] == 10; i-- {
        digits[i] = 0
        digits[i-1]++
    }

    // Handle overflow at most significant digit
    if digits[0] == 10 {
        digits[0] = 1
        digits = append(digits, 0)
    }

    return digits
}

func main() {
    digits := []int{9, 9, 9, 9, 9, 9, 9}
    result := PlusOne(digits)
    fmt.Printf("This is the final result: %v\n", result)
}

