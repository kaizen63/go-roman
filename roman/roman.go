// Package roman provides a converter for integer to roman numbers
package roman

import (
	"strings"
)

// Check if a string contains only roman characters
func IsRoman(r string) bool {
	rl := []string{"I", "V", "X", "L", "C", "D", "M", "ↁ", "ↁ"}
	for _, c := range rl {
		r = strings.Replace(r, c, "", -1)
	}
	if len(r) > 0 {
		return false
	}
	return true
}

// Roman converts an integer to a roman string
func Roman(n int) string {
	if n < 1 {
		return "ROMAN_OUT_OF_RANGE"
	}
	// +------+-----+-----+-----+-----+-----+-----+-----+-----+------+
	// | Char |  I  |  V  |  X  |  L  |  C  |  D  |  M  |  ↁ  |  ↂ  |
	// | Value|  1  |  5  | 10  | 50  | 100 | 500 | 1000| 5000| 10000|
	// +------+-----+-----+-----+-----+-----+-----+-----+-----+------+
	var transTab = []struct {
		old string
		new string
	}{
		{"IIIII", "V"},
		{"IIII", "IV"}, // subtraction rule
		{"VV", "X"},
		{"VIV", "IX"},
		{"XXXXX", "L"},
		{"XXXX", "XL"},
		{"LL", "C"},
		{"LXL", "XC"},
		{"CCCCC", "D"},
		{"CCCC", "CD"},
		{"DD", "M"},
		{"DCD", "CM"},
		{"MMMMM", "ↁ"},
		{"MMMM", "Mↁ"},
		{"ↁↁ", "ↂ"},
	}
	r := strings.Repeat("I", n)
	for _, tt := range transTab {
		r = strings.Replace(r, tt.old, tt.new, -1)
	}
	return r
}

// Convert the argument Roman string to an arabic integer.
func Arabic(roman string) int {
	if !IsRoman(roman) {
		return 0
	}
	r2a := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
		'ↁ': 5000,
		'ↂ': 10000,
	}

	last := 10000
	arabic := 0
	c := []rune(roman)
	for _, chr := range c {
		current := r2a[chr]
		// apply substraction rule if current value gt last. E.g. IX
		if last < current {
			arabic -= 2 * last
		}
		last = current
		arabic += current
	}
	return arabic
}
