package main

import "fmt"

func main() {
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}
}
