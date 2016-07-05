package main

import (
	"bufio"
	"bytes"
	"strings"
	"fmt"
	"net"
	"strconv"
	"github.com/golang/protobuf/proto"
	msg "github.com/jakewins/sickle/gazebo/v7.1/gazebo_msgs"
	"errors"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:11345")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

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
		fmt.Println(publisher.GetHost() + ":" +
			strconv.Itoa(int(publisher.GetPort())) +
			"\t" + publisher.GetTopic() +
			"\t" + publisher.GetMsgType())
	}

	server, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(server.Addr().String(), ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		panic(err)
	}

	advert := &msg.Publish{
		Topic: proto.String("/gazebo/default/Sickle::link_2_JOINT_4/joint_cmd"),
		MsgType: proto.String("gazebo.msgs.JointCmd"),
		Host: proto.String("127.0.0.1"),
		Port: proto.Uint32(uint32(port)),
	}

	if err = writeMessage(writer, "advertise", advert); err != nil {
		panic(err)
	}
	if err = writer.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("Waiting for connection..")
	reverseConn, err := server.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Println("Got a caller!")
	reverseReader := bufio.NewReader(reverseConn)
	for {
		msgtype, _, err := readMessage(reverseReader)
		if err != nil {
			panic(err)
		}
		fmt.Println(msgtype)
	}
}

func writeMessage(writer *bufio.Writer, msgtype string, message proto.Message) (err error) {
	marshalled, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	timestamp := time.Now()
	marshalled, err = proto.Marshal(&msg.Packet{
		Stamp: &msg.Time{
			Sec: proto.Int32(int32(timestamp.Unix())),
			Nsec: proto.Int32(int32(timestamp.Nanosecond())),
		},
		Type: proto.String(msgtype),
		SerializedData:marshalled,
	})
	if err != nil {
		return err
	}

	// TODO fmt should be able to do this with %08x
	packetLen := lpad(strconv.FormatInt(int64(len(marshalled)), 16), "0", 8)
	_, err = writer.Write([]byte(packetLen))
	if err != nil {
		// written == len(packetLen) | err == nil, so no looping needed
		return err
	}

	_, err = writer.Write(marshalled)
	if err != nil {
		return err
	}

	return nil
}

func readMessageOfType(reader *bufio.Reader, assertType string) (message interface{}) {
	gotType, message, err := readMessage(reader)
	if err != nil {
		panic(err)
	}
	if gotType != assertType {
		panic(errors.New("Expected a '" + assertType + "' message, but got '"+gotType+"'"))
	}
	return message
}

func readMessage(reader *bufio.Reader) (msgtype string, message interface{}, err error) {
	packet, err := readPacket(reader)
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

func readPacket(reader *bufio.Reader) (packet *msg.Packet, err error) {
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

func lpad(s string, padStr string, targetLen int) string {
	padLen := 1 + ((targetLen - len(padStr)) / len(padStr))
	result := strings.Repeat(padStr, padLen) + s
	return result[(len(result) - targetLen):]
}