package service

import (
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	f := Feature{
		Cmd: "ls",
	}
	called := false
	cmdStart = func(*exec.Cmd) error { called = true; return nil }
	f.Start()
	if !called {
		t.Errorf("command didn't start")
	}
}
