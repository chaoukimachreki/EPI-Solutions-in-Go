package main

import (
    "math/rand"
    "time"
    "fmt"
)

// RandomSubset shuffles in-place so that A[:k] holds a uniformly random subset.
func RandomSubset[A any](Arr []A, k int) {
    rand.Seed(time.Now().UnixNano()) // Ensure randomness per run

    n := len(Arr)
    for i := 0; i < k; i++ {
        // Pick random index from i to n-1
        r := rand.Intn(n - i) + i
        Arr[i], Arr[r] = Arr[r], Arr[i]
    }
    // A[:k] now holds the random subset
}

func main() {
    customers := []string{"Lina", "Ahmed", "Sara", "John", "Zayd", "Emi"}
    k := 4

    RandomSubset(customers, k)
    fmt.Println("Random subset:", customers[:k])
}

