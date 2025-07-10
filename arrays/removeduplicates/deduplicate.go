package main

import "fmt"

func Deduplicate(origin []int) []int {
    if len(origin) == 0 {
        return origin
    }

    toErase := 1
    for i := 1; i < len(origin); i++ {
        if origin[i] != origin[toErase-1] {
            origin[toErase] = origin[i]
            toErase++
        }
    }
    return origin[:toErase]
}

func main() {
    sample := []int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 6, 6}
    result := Deduplicate(sample)
    fmt.Printf("🧹 Cleaned array: %v\n", result)
}

