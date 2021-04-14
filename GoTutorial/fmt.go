package main

import (
	"fmt"
	"io"
	"os"
)

//func Sprintf(format string, a ...interface{}) string

func main() {
	const name, dept = "geeksforgeeks", "CS"

	// calling Sprinf function
	s := fmt.Sprintf("%s is a %s portal.\n", name, dept)

	io.WriteString(os.Stdout, s)

}
