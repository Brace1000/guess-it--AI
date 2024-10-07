package main

import (
	// "fmt"
	// "os"

	"fmt"
	"os"

	"guess/readdata"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	readdata.ReadData()
}
