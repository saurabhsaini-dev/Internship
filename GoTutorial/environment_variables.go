package main

import (
	"fmt"
	"os"
)

func main() {
	// Set an Environment Variable
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_NAME", "test")

	// Get the value of an Environment Variable
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	fmt.Printf("Host: %s, Port: %s\n", host, port)

	// Unset an Environment Variable
	os.Unsetenv("DB_PORT")
	fmt.Printf("After unset, Port: %s\n", os.Getenv("DB_PORT"))

	/*
		Get the value of an environment variable and a boolean indicating whether the
		environment variable is present or not.
	*/
	driver, ok := os.LookupEnv("DB_DRIVER")
	if !ok {
		fmt.Println("DB_DRIVER is not present")
	} else {
		fmt.Printf("Driver: %s\n", driver)
	}

	// Expand a string containing environment variables in the form of $var or ${var}
	dbURL := os.ExpandEnv("postgres://$DB_USERNAME:$DB_PASSWORD@DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println("DB URL: ", dbURL)

}
