package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMainForMissingArgument(t *testing.T) {
	if os.Getenv("DO_CRASH") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainForMissingArgument")
	cmd.Env = append(os.Environ(), "DO_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMainForWrongArgument(t *testing.T) {
	if os.Getenv("DO_CRASH") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=estMainForWrongArgument", "ABC")
	cmd.Env = append(os.Environ(), "DO_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMain(t *testing.T) {
	os.Args = append(os.Args, "10")
	main()
}
