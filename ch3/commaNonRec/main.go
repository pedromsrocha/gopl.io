// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/commaNonRec.go
//	$ ./main 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	j := len(s) % 3
	if j == 0 {
		j = 3
	}
	for i, v := range s {
		if i == j {
			buf.WriteByte(',')
			j = j + 3
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
