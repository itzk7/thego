// qn: print index, val of the cli arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	osArguments := os.Args

	for index, value := range osArguments {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
