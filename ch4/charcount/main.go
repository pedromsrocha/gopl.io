// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not enough arguments\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "charcount":
		counts := make(map[rune]int)    // counts of Unicode characters
		var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
		invalid := 0                    // count of invalid UTF-8 characters
		letters := 0
		digits := 0

		in := bufio.NewReader(os.Stdin)
		for {
			r, n, err := in.ReadRune() // returns rune, nbytes, error
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
				os.Exit(1)
			}
			if r == unicode.ReplacementChar && n == 1 {
				invalid++
				continue
			}

			if unicode.IsLetter(r) {
				letters++
			}

			if unicode.IsDigit(r) {
				digits++
			}

			counts[r]++
			utflen[n]++
		}

		fmt.Printf("rune\tcount\n")
		for c, n := range counts {
			fmt.Printf("%q\t%d\n", c, n)
		}
		fmt.Print("\nlen\tcount\n")
		for i, n := range utflen {
			if i > 0 {
				fmt.Printf("%d\t%d\n", i, n)
			}
		}
		if invalid > 0 {
			fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
		}
		if letters > 0 {
			fmt.Printf("\n%d letters\n", letters)
		}
		if digits > 0 {
			fmt.Printf("\n%d digits\n", digits)
		}
	case "wordfreq":
		wordCount := map[string]int{}
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			wordCount[input.Text()]++
		}
		fmt.Printf("\nword\tcount\n")
		for word, count := range wordCount {
			fmt.Printf("%q\t%d\n", word, count)
		}

	default:
		fmt.Fprintf(os.Stderr, "Argument is not valid\n")
		os.Exit(1)

	}
}

//!-
