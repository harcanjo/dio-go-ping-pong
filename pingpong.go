package main

import (
	"fmt"
	"time"
)

// pingar sends "ping" messages to the channel c
func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // send "ping" to the channel
	}
}

// pongar sends "pong" messages to the channel c
func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // send "pong" to the channel
	}
}

// imprimir receives messages from the channel c and prints them to the console
func imprimir(c chan string) {
	for {
		msg := <-c // receive a message from the channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // wait for one second before printing the next message
	}
}

func main() {
	var c chan string = make(chan string) // create a new channel of type string
	go pingar(c)                          // start a new goroutine running the pingar function
	go imprimir(c)                        // start a new goroutine running the imprimir function
	go pongar(c)                          // start a new goroutine running the pongar function
	var entrada string
	fmt.Scanln(&entrada) // wait for user input before exiting
}