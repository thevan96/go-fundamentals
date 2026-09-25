package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
		if scanner.Err() != nil {
			break
		}
	}

	for i := range nums {
		increment(&nums[i])
	}

	for _, num := range nums {
		fmt.Println(num)
	}
}

func increment(n *int) {
	*n++
}
