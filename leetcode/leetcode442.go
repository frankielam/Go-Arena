package main

import (
	"fmt"
)

func findDuplicates(nums []int) (ans []int) {
	for i := range nums {
        for nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        } 
    }
    for i, num := range nums {
        if (num - 1 != i) {
            ans = append(ans, num)
        }
    }
    return ans
}

func main() {
	arr := []int {4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(arr))
}