package main

import (
	"fmt"
)

func main() {
	frequencies := map[string]int{}
	var word string

	for {
		_, err := fmt.Scan(&word)
		if err != nil {
			break
		}
		frequencies[word]++
	}

	different := len(frequencies)
	once := 0
	for _, v := range frequencies {
		if v == 1 {
			once++
		}
	}
	fmt.Println(different)
	fmt.Println(once)
}
