package main

import (
	"fmt"

	"github.com/akakou/crtsh"
)

func main() {
	data, err := crtsh.Fetch("test.ochano.co")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Doamin[0]: %v\n", data[0].NameValue)
	fmt.Printf("Index[0]: %v\n", data[0].ID)
}
