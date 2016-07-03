package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"math"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:11345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)
	for {
		// Header: 8 bytes, a hex string denoting the length of the packet
		str, err := string(readChunk(reader, 8))
		if err != nil {
			break // TODO
		}
		packetLen, err := strconv.ParseInt(str, 16, 32)
		if err != nil {
			break // TODO
		}
		packet, err := readChunk(packetLen)
		if err != nil {
			break // TODO
		}

		break
	}
}

func readChunk(reader *bufio.Reader, len int) (bytes.Buffer, error) {
	var buf [1024]byte
	var out bytes.Buffer
	var remaining = len

	for remaining > 0 {
		read, err := reader.Read(buf[:math.min(len(buf), remaining)])
		remaining -= read
		if err != nil {
			return nil, err
		}
		if read > 0 {
			out.Write(buf[:read])
		}
	}
	return out, nil
}
