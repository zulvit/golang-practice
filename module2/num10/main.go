package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	// Запуск go-рутины для чтения из сокета и записи в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from socket: %v\n", err)
			os.Exit(1)
		}
	}()

	// Чтение из STDIN и запись в сокет
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to socket: %v\n", err)
		os.Exit(1)
	}
}
