package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const Healthy = 0
const Unhealthy = 1
const MaximumStringSize = 4096 // 4 KiB

func main() {
	address := flag.String("a", "127.0.0.1:4325", "address")
	flag.Parse()

	conn, err := net.Dial("tcp", *address)
	if err != nil {
		fmt.Println("dial error:", err)
		os.Exit(3)
	}

	lr := io.LimitReader(conn, 1+MaximumStringSize)
	buf, err := io.ReadAll(lr)
	if err != nil {
		fmt.Println("read error:", err)
		os.Exit(3)
	}

	err = conn.Close()
	if err != nil {
		fmt.Println("close error:", err)
	}

	var state bool
	var message string
	switch len(buf) {
	case 0:
		state = false
	case 1:
		state = buf[0] == Healthy
	default:
		state = buf[0] == Healthy
		message = string(buf[1:])
	}

	fmt.Print(message)
	if state {
		os.Exit(Healthy)
	} else {
		os.Exit(Unhealthy)
	}
}
