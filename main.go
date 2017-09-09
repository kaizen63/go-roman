package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	"github.com/kaizen63/go-roman/roman"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	args := flag.Args()
	var max int
	var err error
	if len(args) > 0 {
		max, err = strconv.Atoi(flag.Args()[0])
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

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}
