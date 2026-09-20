package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if scanner.Err() != nil {
			os.Exit(1)
		}

		return
	}

	numRaw := scanner.Text()
	if num, err := strconv.Atoi(numRaw); err != nil {
		fmt.Println("bad")
		return
	} else {
		fmt.Printf("ok %d\n", num)
	}
}
