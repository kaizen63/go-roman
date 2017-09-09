[![Build Status](https://travis-ci.org/kaizen63/go-roman.svg?branch=master)](https://travis-ci.org/kaizen63/go-roman)

# go-roman

A small package to convert integers to roman numbers. The intersting bit is that the conversion into roman is not a mathematical solution. Its pure text processing.
Kudos to Kevlin Hennley for this solution (although it was not in go)
## Example
```
package main

import (
	"fmt"
	"os"
	"strconv"

	roman "github.com/kaizen63/go-roman"
)

func main() {
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for i := 1; i <= max; i++ {
		fmt.Printf("%4d : %s\n", i, roman.RomanValueOf(i))
	}
}
```