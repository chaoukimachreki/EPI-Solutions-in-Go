package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

// OnlineRandomSample reads from a stream (stdin here) and returns a random subset of size k.
func OnlineRandomSample(k int) []int {
    rand.Seed(time.Now().UnixNano())

    sample := make([]int, 0, k)
    scanner := bufio.NewScanner(os.Stdin)

    numSeen := 0

    // Read integers from stdin line-by-line
    for scanner.Scan() {
        text := scanner.Text()
        x, err := strconv.Atoi(text)
        if err != nil {
            continue // Skip non-integer input
        }
        numSeen++

        if len(sample) < k {
            sample = append(sample, x)
        } else {
            // Generate a random index in [0, numSeen-1]
            idx := rand.Intn(numSeen)
            if idx < k {
                sample[idx] = x
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading stream:", err)
    }
    return sample
}

func main() {
    k := 4
    fmt.Printf("Enter stream of integers (Ctrl+D or EOF to end):\n")
    sample := OnlineRandomSample(k)
    fmt.Println("Random sample:", sample)
}

