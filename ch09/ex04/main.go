package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var num int
	var head, tail chan struct{}
	var start time.Time

	num = 1e3
	head, tail = makeLinkedChannel(num)
	start = time.Now()
	head <- struct{}{}
	<-tail
	fmt.Printf("number of channels: %d, %s\n", num, time.Since(start))

	num = 1e5
	head, tail = makeLinkedChannel(num)
	start = time.Now()
	head <- struct{}{}
	<-tail
	fmt.Printf("number of channels: %d, %s\n", num, time.Since(start))

	num = 1e6
	head, tail = makeLinkedChannel(num)
	start = time.Now()
	head <- struct{}{}
	<-tail
	fmt.Printf("number of channels: %d, %s\n", num, time.Since(start))

	num = math.MaxInt32
	head, tail = makeLinkedChannel(num)
	start = time.Now()
	head <- struct{}{}
	<-tail
	fmt.Printf("number of channels: %d, %s\n", num, time.Since(start))
}

func makeLinkedChannel(n int) (head, tail chan struct{}) {
	head = make(chan struct{})
	tail = head
	for i := 0; i < n-1; i++ {
		tail = appendlinkedChannel(tail)
	}
	return
}

func appendlinkedChannel(prev chan struct{}) chan struct{} {
	next := make(chan struct{})
	go func() {
		val := <-prev
		next <- val
	}()
	return next
}
