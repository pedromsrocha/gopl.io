// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	s2 := [5]int{1, 2, 3, 4, 5}
	reverse2(&s2)
	fmt.Println(s2)

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		input.Scan()
		opt := strings.Fields(input.Text())
		switch opt[0] {
		case "rev":
			reverse(ints)
			fmt.Printf("%v\n", ints)
		case "rot":
			x, err := strconv.ParseInt(opt[1], 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = rotate(ints, int(x))
			fmt.Printf("%v\n", ints)
		default:
			continue outer
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev

// ex4.3: same as reverse above, but since we are using an array pointer we need to fix the length
const N = 5

func reverse2(s *[N]int) {
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate s left by n positions
func rotate(s []int, n int) []int {
	N := len(s)
	srot := make([]int, N)
	for i := 0; i <= N-1; i++ {
		srot[i] = s[(i+n)%N]
	}
	return srot
}
