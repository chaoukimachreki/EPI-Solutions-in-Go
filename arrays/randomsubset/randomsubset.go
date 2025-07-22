package main

import (
    "fmt"
    "math/rand"
    "time"
)

// RandomSubset returns a slice containing a random k-sized subset of {0, 1, ..., n-1}
func RandomSubset(n, k int) []int {
    changed := make(map[int]int)
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < k; i++ {
        randIdx := rand.Intn(n-i) + i

        valI, okI := changed[i]
        valRandIdx, okRand := changed[randIdx]

        switch {
        case !okI && !okRand:
            changed[i], changed[randIdx] = randIdx, i
        case !okI && okRand:
            changed[i] = valRandIdx
            changed[randIdx] = i
        case okI && !okRand:
            changed[randIdx] = valI
            changed[i] = randIdx
        default:
            changed[randIdx] = valI
            changed[i] = valRandIdx
        }
    }

    result := make([]int, k)
    for i := 0; i < k; i++ {
        if val, ok := changed[i]; ok {
            result[i] = val
        } else {
            result[i] = i
        }
    }

    return result
}

func main() {
    n, k := 10, 4
    subset := RandomSubset(n, k)
    fmt.Println("Random subset:", subset)
}

