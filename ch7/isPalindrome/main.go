package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	i := 0
	j := s.Len() - 1
	for i < j {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	var word sort.IntSlice

	word = sort.IntSlice([]int{1, 2, 3})
	fmt.Printf("Is %v a palindrome? %t\n", word, IsPalindrome(word))

	word = sort.IntSlice([]int{1, 2, 3, 2, 1})
	fmt.Printf("Is %v a palindrome? %t\n", word, IsPalindrome(word))

	word = sort.IntSlice([]int{4, 2, 4})
	fmt.Printf("Is %v a palindrome? %t\n", word, IsPalindrome(word))

	word = sort.IntSlice([]int{1, 0, 0, 1})
	fmt.Printf("Is %v a palindrome? %t\n", word, IsPalindrome(word))

	word = sort.IntSlice([]int{1, 0, 1, 0})
	fmt.Printf("Is %v a palindrome? %t\n", word, IsPalindrome(word))
}
