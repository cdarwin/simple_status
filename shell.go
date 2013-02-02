package main

import (
	"log"
	"os/exec"
)

type Shell struct {
	Output string `json:"out"`
}

func shell(s string) Shell {
	out, err := exec.Command("/bin/sh", "-c", s).Output()
	if err != nil {
		log.Println(err)
	}
	var sh Shell
	sh.Output = string(out)
	return sh
}
