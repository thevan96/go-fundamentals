package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
