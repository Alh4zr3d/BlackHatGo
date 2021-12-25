package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		var address = fmt.Sprintf("scanme.nmap.org:%d", p)
		var conn, err = net.Dial("tcp", address)
		if err == nil {
			results <- p
			conn.Close()
		} else {
			results <- 0
		}
	}
}

func main() {
	var ports = make(chan int, 100)
	var results = make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
