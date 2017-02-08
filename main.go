package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {

	fmt.Println("Socket receiver")

	fmt.Println("Start TCP server")

	// listen on a port
	ln, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		fmt.Println("Unable to listen on TCP port")
		return
	}

	var count = 0
	var wg sync.WaitGroup

	// wait for a connection
	fmt.Println("Server waiting for a connection")
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting connection")
		return
	}
	fmt.Println("Server: Connection accepted!")

	// begin reading from connection
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Empty message")
			break
		}

		count++

		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Println("All go routines done")
	conn.Close()
	fmt.Println(count, " messages received")
}
