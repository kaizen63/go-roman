package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/kaizen63/go-roman/roman"
)

func parseCommandline() (max int, err error) {
	flag.Parse()
	args := flag.Args()
	var e error
	if len(args) > 0 {
		max, e = strconv.Atoi(args[0])
		if e != nil {
			fmt.Println(err)
			return 0, fmt.Errorf("wrong argument, got '%s' expecting integer value", args[0])
		}
	} else {
		fmt.Println("usage: go-roman <max>")
		return 0, fmt.Errorf("usae: go-roman <max>")
	}
	return max, nil
}

func main() {
	max, err := parseCommandline()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= max; i++ {
		fmt.Println(roman.Roman(i))
	}

}
