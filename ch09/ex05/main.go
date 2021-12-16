package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	result := make(chan int)

	go pingpong(done, result)
	<-time.After(1 * time.Second)
	close(done)

	res := <-result

	fmt.Printf("%d rallies in a second.\n", res)
}

func pingpong(done <-chan struct{}, result chan<- int) {
	var rally int
	ping := make(chan struct{})
	pong := make(chan struct{})

	go func() { ping <- struct{}{} }()

	go func() {
		for {
			select {
			case val := <-ping:
				rally++
				pong <- val
			case <-done:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case val := <-pong:
				rally++
				ping <- val
			case <-done:
				return
			}
		}
	}()

	<-done
	result <- rally
}
