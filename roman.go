package roman

import (
	"strings"
)

// RomanValueOf converts an integer to a roman string
// +------+-----+-----+-----+-----+-----+-----+-----+-----+------+
// | Char |  I  |  V  |  X  |  L  |  C  |  D  |  M  |  ↁ  |  ↂ  |
// | Value|  1  |  5  | 10  | 50  | 100 | 500 | 1000| 5000| 10000|
// +------+-----+-----+-----+-----+-----+-----+-----+-----+------+
func RomanValueOf(n int) string {
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
	if n < 1 {
		return "ROMAN_OUT_OF_RANGE"
	}

	r := strings.Repeat("I", n)

	for _, tt := range transTab {
		r = strings.Replace(r, tt.old, tt.new, -1)
	}
	return r
}
