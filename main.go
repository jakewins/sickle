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
	"errors"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:11345")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)

	init := readMessageOfType(reader, "version_init").(*msg.GzString)
	if(init.GetData() != "gazebo 7.2") {
		panic("Unknown Gazebo version: " + init.GetData())
	}

	namespaces := readMessageOfType(reader, "topic_namepaces_init").(*msg.GzString_V).GetData()
	for _, namespace := range namespaces {
		fmt.Println(namespace)
	}

	publishers := readMessageOfType(reader, "publishers_init").(*msg.Publishers)
	for _, publisher := range publishers.GetPublisher() {
		fmt.Println(publisher.GetHost() + ":" + strconv.Itoa(int(publisher.GetPort())) + "\t" + publisher.GetTopic())
	}
}

func readMessageOfType(reader *bufio.Reader, assertType string) (message interface{}) {
	gotType, message, err := readGazeboMessage(reader)
	if err != nil {
		panic(err)
	}
	if gotType != assertType {
		panic(errors.New("Expected a '" + assertType + "' message, but got '"+gotType+"'"))
	}
	return message
}

func readGazeboMessage(reader *bufio.Reader) (msgtype string, message interface{}, err error) {
	packet, err := readGazeboPacket(reader)
	if err != nil {
		return "", nil, err
	}
	msgtype = packet.GetType()

	switch msgtype {
	case "version_init":
		message = &msg.GzString{}
	case "topic_namepaces_init":
		message = &msg.GzString_V{}
	case "publishers_init":
		message = &msg.Publishers{}
	default:
		return msgtype, nil, errors.New("Unknown message type: " + packet.GetType())
	}
	err = proto.Unmarshal(packet.GetSerializedData(), message.(proto.Message))
	if err != nil {
		return msgtype, nil, err
	}

	return msgtype, message, nil
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