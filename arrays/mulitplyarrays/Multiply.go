package main

import (
    "fmt"
    "errors"
)

// Multiply multiplies two numbers represented as slices of digits (base-10).
func Multiply(num1, num2 []int) ([]int, error) {
    if len(num1) == 0 || len(num2) == 0 {
        return nil, errors.New("input slices must be non-empty")
    }

    sign := 1
    if (num1[0] < 0) != (num2[0] < 0) {
        sign = -1
    }

    // Convert to absolute values
    for i := range num1 {
        if num1[i] < 0 {
            num1[i] = -num1[i]
        }
    }
    for i := range num2 {
        if num2[i] < 0 {
            num2[i] = -num2[i]
        }
    }

    result := make([]int, len(num1)+len(num2))

    // Multiply digits from right to left
    for i := len(num1) - 1; i >= 0; i-- {
        for j := len(num2) - 1; j >= 0; j-- {
            product := num1[i] * num2[j]
            sum := result[i+j+1] + product
            result[i+j+1] = sum % 10
            result[i+j] += sum / 10
        }
    }

    // Trim leading zeros
    start := 0
    for start < len(result)-1 && result[start] == 0 {
        start++
    }
    result = result[start:]

    // Apply sign
    if sign == -1 {
        result[0] = -result[0]
    }

    return result, nil
}

func main() {
    num1 := []int{1, 2, 3}    // Represents 123
    num2 := []int{-4, 5, 6}   // Represents -456

    result, err := Multiply(num1, num2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("The result of the multiplication is: %v\n", result)
}

