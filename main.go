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
		packet, nil := readGazeboPacket(reader)
		if err != nil {
			break // TODO
		}

		fmt.Println(packet.GetType())
		break
	}
}

func readGazeboPacket(reader *bufio.Reader) (packet *msg.Packet, err error) {
	// Header: 8 byte hex string denoting length of payload
	chunk, err := readChunk(reader, 8)
	if err != nil {
		return nil, err
	}
	packetLen, err := strconv.ParseInt(chunk.String(), 16, 32)
	if err != nil {
		return nil, err
	}

	// Packet body
	chunk, err = readChunk(reader, int(packetLen))
	if err != nil {
		return nil, err
	}
	packet = &msg.Packet{}
	err = proto.Unmarshal(chunk.Bytes(), packet)
	if err != nil {
		return nil, err
	}

	return packet, nil
}

func readChunk(reader *bufio.Reader, chunkSize int) (chunk *bytes.Buffer, err error) {
	var buf [1024]byte
	var out bytes.Buffer
	var remaining = chunkSize

	for remaining > 0 {
		read, err := reader.Read(buf[:min(len(buf), remaining)])
		remaining -= read
		if err != nil {
			return nil, err
		}
		if read > 0 {
			out.Write(buf[:read])
		}
	}
	return &out, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}