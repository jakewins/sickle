package main

import (
	"bufio"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(conn)

	writer.Write([]byte("1 atrv.motion set_speed [1,1]\n"))
	writer.Flush()
}