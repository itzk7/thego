// qn: print the file which contains duplicate values

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filesCountMapper := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, filesCountMapper)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, filesCountMapper)
			f.Close()
		}
	}
	for fileName, counts := range filesCountMapper {
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, fileName)
			}
		}	
	}
	
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}