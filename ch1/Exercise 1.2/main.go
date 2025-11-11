package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[:] {
		if i == 0 {
			continue // Don't print arg 0, the name of the command.
		}
		fmt.Println("Argument " + strconv.Itoa(i) + ": " + arg) // Alternatively, i + 1 here.
	}
}
