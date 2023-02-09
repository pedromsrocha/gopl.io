// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	data = []string{"A", "B", "B", "B", "C", "C", "D", "D", "D", "D"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", remdups(data))
	//!-main
}

// !+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt

// ex 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice
// notice: the underlying array of strings is modified during the call
func remdups(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	prev := strings[0]
	i := 1
	for _, s := range strings {
		if s != prev {
			strings[i] = s
			prev = s
			i++
		}
	}
	return strings[:i]
}
