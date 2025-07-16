package main

import (
    "fmt"
    "math/rand"
    "time"
)

// ComputeRandomPermutation returns a random permutation of [0, 1, ..., n-1]
func ComputeRandomPermutation(n int) []int {
    rand.Seed(time.Now().UnixNano())

    perm := make([]int, n)
    for i := range perm {
        perm[i] = i
    }

    // Fisher–Yates shuffle
    for i := 0; i < n; i++ {
        r := rand.Intn(n - i) + i
        perm[i], perm[r] = perm[r], perm[i]
    }

    return perm
}

func main() {
    n := 3
    fmt.Println("Random permutation:", ComputeRandomPermutation(n))
}

