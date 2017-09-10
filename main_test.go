package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMainForMissingArgumentExec(t *testing.T) {
	if os.Getenv("DO_CRASH") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainForMissingArgumentExec")
	cmd.Env = append(os.Environ(), "DO_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMainForWrongArgumentExec(t *testing.T) {
	if os.Getenv("DO_CRASH") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainForWrongArgumentExec", "ABC")
	cmd.Env = append(os.Environ(), "DO_CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestMain(t *testing.T) {
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	os.Args = append(os.Args, "10")
	main()
	os.Stdout = old
}

func TestParseCommandLine(t *testing.T) {
	os.Args = []string{""}
	os.Args = append(os.Args, "123")
	v, err := parseCommandline()
	if err != nil {
		t.Errorf("Expecting success, got error")
	}
	if v != 123 {
		t.Errorf("Expected 10 got %v", v)
	}
}
func TestForMissingArgument(t *testing.T) {
	os.Args = []string{""}
	_, err := parseCommandline()
	if err == nil {
		t.Errorf("Expecting missing parameter error. Got success")
	}
}

func TestForWrongArgument(t *testing.T) {
	os.Args = []string{""}
	os.Args = append(os.Args, "ABC")
	_, err := parseCommandline()
	if err == nil {
		t.Errorf("Expecting wrong parameter error. Got success")
	}
}
