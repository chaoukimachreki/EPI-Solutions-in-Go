package main

import("fmt")


func NextPermutation(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }

    // Step 1: Find the pivot (first index from right where nums[i] < nums[i+1])
    k := n - 2
    for k >= 0 && nums[k] >= nums[k+1] {
        k--
    }

    // If no pivot is found, we're at the last permutation
    if k < 0 {
        return false
    }

    // Step 2: Find the smallest number greater than nums[k] to the right
    l := n - 1
    for nums[l] <= nums[k] {
        l--
    }

    // Step 3: Swap pivot and successor
    nums[k], nums[l] = nums[l], nums[k]

    // Step 4: Reverse the suffix starting at k+1
    for left, right := k+1, n-1; left < right; left, right = left+1, right-1 {
        nums[left], nums[right] = nums[right], nums[left]
    }

    return true
}

func main(){

        perm := []int{1,7,8,9,3,4,5,6,2}
	if NextPermutation(perm) {
		fmt.Println("Next permutation:", perm) // Output: [1 2 0 3]
	} else {
		fmt.Println("Already at last permutation")
	}



}
