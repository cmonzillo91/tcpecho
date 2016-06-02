package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8181")
	if err != nil {
		fmt.Printf("Could not start listener: %s\n", err)
		return
	}
	for {
		fmt.Println("Listening for Connections")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Could not accept connection!!!")
			continue
		}
		go func() {
			fmt.Println("Creating reader and writer for the connection")
			r, w := bufio.NewReader(conn), bufio.NewWriter(conn)

			for {
				rs, err := r.ReadString('\n')
				if err != nil {
					fmt.Printf("We OUT: %s\n", err)
					break
				}
				fmt.Print(rs)
				w.WriteString(rs)
				w.Flush()

			}
		}()
	}
}
