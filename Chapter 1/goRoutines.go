package main

import (
	"fmt"
)

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	var c = make(chan int)
	go strlen("Cthulhu", c)
	go strlen("Nyarlathotep", c)
	var x, y = <-c, <-c
	fmt.Println(x, y, x+y)
}
