package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.Copy(conn, conn)
		if err != nil {
			log.Fatalln("[-] Error reading/writing data.")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":18110")
	if err != nil {
		log.Fatalln("[-] Unable to bind to port.")
	}
	log.Println("[+] Listening on 0.0.0.0:18110...")
	for {
		conn, err := listener.Accept()
		log.Println("[+] Connection received.")
		if err != nil {
			log.Fatalln("[-] Unable to accept connection.")
		}
		go echo(conn)
	}
}
