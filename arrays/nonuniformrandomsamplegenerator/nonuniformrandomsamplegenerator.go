package main

import (
    "fmt"
    
    "math/rand"
    "sort"
    "time"
)

// NonuniformRandomNumberGeneration returns a random value from 'values'
// based on the corresponding probabilities.
func NonuniformRandomNumberGeneration(values []int, probabilities []float64) int {
    prefixSums := make([]float64, len(probabilities)+1)
    prefixSums[0] = 0.0

    for i := range probabilities {
        prefixSums[i+1] = prefixSums[i] + probabilities[i]
    }

    rand.Seed(time.Now().UnixNano())
    uniform01 := rand.Float64()

    // Find the right interval via binary search
    intervalIdx := sort.Search(len(prefixSums), func(i int) bool {
        return prefixSums[i] > uniform01
    }) - 1

    return values[intervalIdx]
}

func main() {
    values := []int{3, 5, 7, 11}
    probs := []float64{0.1, 0.3, 0.4, 0.2}
    fmt.Println("Random value:", NonuniformRandomNumberGeneration(values, probs))
}

