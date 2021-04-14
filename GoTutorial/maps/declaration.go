package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// Attempting to add keys to a nil map will result in a runtime error
	// m["one hundred"] = 100
}
