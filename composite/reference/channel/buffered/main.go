package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "a"
	ch <- "b"
	ch <- "c"

	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
