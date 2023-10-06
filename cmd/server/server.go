package main

import (
	"fmt"
	"net"
	"os"
)

const port = ":3000"

func main() {
	listenAddress, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot resolve TCP address %q", err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp4", listenAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot listen on port %ss %q", port, err)
		os.Exit(1)
	}

	fmt.Printf("🥝 server is listening on %s\n", listenAddress.String())

	for {
		con, err := listener.AcceptTCP()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot accept incoming TCP connection request %q", err)
		}

		go handleClient(con)
	}
}

func handleClient(con *net.TCPConn) {
	const maxReadSize = 512
	buf := make([]byte, 512)
	for {
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println("😴 client has disconnected")
			return
		}

		fmt.Printf("🧑🏻‍💻 server ↳\t%s\n", buf[:n])
		con.Write(buf[:n])
	}
}
