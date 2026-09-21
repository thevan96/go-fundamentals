package main

import "fmt"

func main() {
	var nums []int
	var x int

	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		nums = append(nums, x)
	}

	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	fmt.Println(maxNum)
}
