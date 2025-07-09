package main

import (
    "errors"
    "fmt"
    "math"
)

// Reverse takes an integer and returns its digits reversed.
// It handles negatives and detects overflow.
func Reverse(x int) (int, error) {
    if x == 0 {
        return 0 , errors.New("are you kidding me , want to reverse 0 ?")
    }

    negative := x < 0
    xAbs := abs(x)
    result := 0

    for xAbs > 0 {
        digit := xAbs % 10

        // Overflow detection before multiplying
        if result > (math.MaxInt32-digit)/10 {
            return 0, errors.New("integer overflow during reversal")
        }

        result = result*10 + digit
        xAbs /= 10
    }

    if negative {
        result = -result
    }

    return result, nil
}

// abs returns the absolute value of an integer.
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
    testCases := []int{
        12345,
        -9876,
        0,
        1000,
        1534236469, // will cause overflow
    }

    for _, x := range testCases {
        result, err := Reverse(x)
        if err != nil {
            fmt.Printf("Reverse(%d): error → %v\n", x, err)
        } else {
            fmt.Printf("Reverse(%d) = %d\n", x, result)
        }
    }
}


