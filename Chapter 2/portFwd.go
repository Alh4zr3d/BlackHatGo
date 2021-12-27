package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "localhost:18111")
	if err != nil {
		log.Fatalln("[-] Unable to connect; target host unreachable.")
	}
	defer dst.Close()

	go forward(src, dst)
	retOutput(src, dst)
}

func forward(src, dst net.Conn) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
}

func retOutput(src, dst net.Conn) {
	_, err := io.Copy(src, dst)
	if err != nil {
		log.Fatalln(err)
	}
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
