package main

import (
    "fmt"
    "math/bits"
)

// ClosestSameWeight finds the closest uint64 to x with the same number of set bits (weight).
// It swaps the rightmost pair of differing adjacent bits.
func ClosestSameWeight(x uint64) (uint64, error) {
    const totalBits = 64

    for i := 0; i < totalBits-1; i++ {
        bitI := (x >> i) & 1
        bitNext := (x >> (i + 1)) & 1

        if bitI != bitNext {
            // Flip the bits at positions i and i+1
            mask := uint64(1)<<i | uint64(1)<<(i+1)
            return x ^ mask, nil
        }
    }

    return 0, fmt.Errorf("no bits to swap in %d: all bits are 0 or 1", x)
}

func main() {

    values:=[]uint64{9,^uint64(0),85,80}

    for _, x := range values {
	    
    	fmt.Printf("Input    : %064b (%d)\n", x, x)

   	closest, err := ClosestSameWeight(x)
        if err != nil {
        	fmt.Println("Error:", err)
        
    	}

    	fmt.Printf("Closest  : %064b (%d)\n", closest, closest)
    	fmt.Printf("Weight   : original=%d, closest=%d\n", bits.OnesCount64(x), bits.OnesCount64(closest))
    }
}




