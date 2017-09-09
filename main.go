package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kaizen63/go-roman/roman"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var max int
	var err error
	if len(args) > 0 {
		max, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	} else {
		fmt.Println("usage: go-roman <max>")
		os.Exit(3)
	}
	for i := 1; i <= max; i++ {
		fmt.Println(roman.RomanValueOf(i))
	}

}
