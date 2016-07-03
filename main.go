package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/golang/protobuf/proto"
	msg "github.com/jakewins/sickle/gazebo/v7.1/gazebo_msgs"
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
		chunk, err := readChunk(reader, 8)
		if err != nil {
			break // TODO
		}
		packetLen, err := strconv.ParseInt(chunk.String(), 16, 32)
		if err != nil {
			break // TODO
		}
		chunk, err = readChunk(reader, int(packetLen))
		if err != nil {
			break // TODO
		}
		packet := &msg.Packet{}
		err = proto.Unmarshal(chunk.Bytes(), packet)
		if err != nil {
			break // TODO
		}

		fmt.Println(packet.Type)
		break
	}
}

func readChunk(reader *bufio.Reader, chunkSize int) (chunk bytes.Buffer, err error) {
	var buf [1024]byte
	var out bytes.Buffer
	var remaining = chunkSize

	for remaining > 0 {
		read, err := reader.Read(buf[:min(len(buf), remaining)])
		remaining -= read
		if err != nil {
			return out, err
		}
		if read > 0 {
			out.Write(buf[:read])
		}
	}
	return out, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}