package main

import "fmt"

// applyPermutation reorders slice A in-place using permutation P.
// Mutates both A and P, avoiding extra allocations.
func applyPermutation(A []string, P []int) {
    for i := 0; i < len(A); i++ {
        next := i
        for P[next] >= 0 {
            A[next], A[P[next]] = A[P[next]], A[next]
            temp := P[next]
            P[next] -= len(A) // mark visited
            next = temp
        }
    }

    // Restore P to original state in-place (no allocation)
    for i := range P {
        P[i] += len(A)
    }
}

func main() {
    A := []string{"a", "b", "c", "d", "e", "f"}
    P := []int{3, 5, 0, 2, 1, 4}
    applyPermutation(A, P)
    fmt.Println("Rearranged:", A)
    fmt.Println("Original P restored:", P)
}

