package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/kaizen63/go-roman/roman"
)

func parseCommandline() (max int, err error) {
	//fmt.Printf("Args: %s\n", strings.Join(os.Args[:], ","))
	flag.Parse()
	args := flag.Args()
	//fmt.Printf("Args: %s\n", strings.Join(args[:], ","))
	var e error
	if len(args) > 0 {
		max, e = strconv.Atoi(args[0])
		if e != nil {
			return 0, fmt.Errorf("wrong argument, got '%s' expecting integer value", args[0])
		}
	} else {
		return 0, fmt.Errorf("usage: go-roman <max>")
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
