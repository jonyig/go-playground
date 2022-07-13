package service

import "os/exec"

var cmdStart = (*exec.Cmd).Start

type Feature struct {
	Cmd string
}

func (f *Feature) Start() error {
	cmd := exec.Command(f.Cmd)
	return cmdStart(cmd)
}
