package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		var reader = bufio.NewReader(conn)
		var writer = bufio.NewWriter(conn)

		s, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln("[-] Unable to read data.")
		}

		log.Printf("[+] Read %d bytes: %s", len(s), s)

		log.Println("[+] Writing data...")

		_, err = writer.WriteString(s)

		if err != nil {
			log.Fatalln("[-] Unable to write data.")
		}

		writer.Flush()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":18110")
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}
	log.Println("Listening on 0.0.0.0:18110...")
	for {
		conn, err := listener.Accept()
		log.Println("Connection received.")
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}
		go echo(conn)
	}
}
