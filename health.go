package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	address := flag.String("a", "127.0.0.1:4325", "address")
	flag.Parse()

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		fmt.Println("dial error:", err)
		os.Exit(3)
	}

	buf := make([]byte, 1)

	n, err := io.ReadFull(conn, buf)
	if err != nil {
		fmt.Println("read error:", err)
		os.Exit(3)
	}

	err = conn.Close()
	if err != nil {
		fmt.Println("close error:", err)
	}

	if n == 1 && buf[0] == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
