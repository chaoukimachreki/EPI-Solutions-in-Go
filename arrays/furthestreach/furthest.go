package main

import (
    "fmt"
)

// CanReachEnd returns true if it's possible to reach the end of the array.
func CanReachEnd(array []int) bool {
    farthest := 0
    for i, length := range array {
        if i > farthest {
            return false // We've hit a point we can't reach.
        }
        farthest = max(farthest, i+length)
    }
    return true
}

// max returns the greater of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    testCases := [][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 0, 5, 6},
        {3, 0, 0, 1, 2},
	{1,0,0,0,9},
	
    }

    for _, arr := range testCases {
        
	    canReach := CanReachEnd(arr)
	    status := "❌  cannot reach end sorry "
	    if canReach {status = "yep ✅ can reach end"}
	    fmt.Printf("%s — Array: %v\n", status, arr)

    }

}

