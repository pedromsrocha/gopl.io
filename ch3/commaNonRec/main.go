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
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	if s != "" && (s[0] == '+' || s[0] == '-') {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	var j int
	k := strings.IndexByte(s, '.')
	if k != -1 {
		j = len(s[:k]) % 3
	} else {
		j = len(s) % 3
		k = len(s)
	}
	if j == 0 {
		j = 3
	}
	for i, _ := range s {
		if i == j && i < k {
			buf.WriteByte(',')
			j = j + 3
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
