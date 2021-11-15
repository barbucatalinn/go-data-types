package main

import "fmt"

func main() {
	ch := make(chan string)
	go addData(ch)

	for event := range ch {
		fmt.Println(event)
	}
}

func addData(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"

	close(ch)
}
