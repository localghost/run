package main

import (
	"os"
	"syscall"
	"log"
	"os/exec"
)

func main() {
	cmd, err := exec.LookPath(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if err := syscall.Exec(cmd, os.Args[1:], os.Environ()); err != nil {
		log.Fatal(err)
	}
}
