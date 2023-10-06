package main

import (
	"fmt"
	"net"
	"os"
)

const serverAddress = "localhost:3000"

func main() {
	server, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot resolve TCP address %q", err)
		os.Exit(1)
	}

	con, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot initiate a connection to server %q", err)
		os.Exit(1)
	}

	fmt.Println("🥑 connected to server ")

	var msg string

	buf := make([]byte, 512)
	go func() {
		for {
			n, err := con.Read(buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot read incoming message %q", err)
				continue
			}

			fmt.Printf("👩🏻‍💻 client ↳\t%s\n", buf[:n])
		}
	}()

	for {
		fmt.Scan(&msg)

		_, err = con.Write([]byte(msg))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot send message to server")
		}
	}
}
