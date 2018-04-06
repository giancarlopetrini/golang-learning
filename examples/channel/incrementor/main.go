package main

import (
	"fmt"
)

func main() {
	c := incrementor(4)
	var count int

	for i := range c {
		count++
		fmt.Println("Process Counter:", count)
		fmt.Println(i)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for num := 0; num < 20; num++ {
				out <- fmt.Sprint("Process: ", i, " printing: ", num)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()

	return out
}
