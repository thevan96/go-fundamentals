package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)
	greetings := fmt.Sprintf("Hi, %s! You are %d years old.", name, age)
	fmt.Println(greetings)
}
