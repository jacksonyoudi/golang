package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	// func NewScanner(r io.Reader) *Scanner {
	//	return &Scanner{
	//		r:            r,
	//		split:        ScanLines,
	//		maxTokenSize: MaxScanTokenSize,
	//	}
	//}
	input := bufio.NewScanner(os.Stdin)

	// ctr + d
	for input.Scan() {
		// input.Text()每一行的输入
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}