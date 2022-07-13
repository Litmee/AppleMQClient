package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
	"time"
)

// Encode message encoding
func Encode(m string) ([]byte, error) {

	//  1. Read the length of the message, convert it to int32 type (4 bytes)
	l := int32(len(m))

	// 2. define an empty bytes buffer
	b := new(bytes.Buffer)

	// 3. Write the message header, and write l to b in a little-endian sequence
	err := binary.Write(b, binary.LittleEndian, l)
	if err != nil {
		return nil, err
	}

	// 4. write message entity
	err = binary.Write(b, binary.LittleEndian, []byte(m))
	if err != nil {
		return nil, err
	}

	// 5. Return the packaged message
	return b.Bytes(), nil
}

type AppleMessage struct {
	// Sign: 0 means from the producer, 1 means from the cluster machine synchronization
	Sign int
	Body string
	Tag  string
}

func doing(c *net.TCPConn) {
	message := &AppleMessage{Sign: 0, Body: time.Now().String(), Tag: "okkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"}

	marshal, _ := json.Marshal(message)

	encode, err := Encode(string(marshal))
	if err != nil {
		log.Println(err)
		return
	}
	_, err = c.Write(encode)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {

	raddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9082")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	encode, _ := Encode("send")
	conn.Write(encode)
	num := 0

	//message := &AppleMessage{Sign: 0, Body: time.Now().String(), Tag: "okkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"}
	//
	//marshal, _ := json.Marshal(message)
	//
	//encode, err = Encode(string(marshal))

	for {
		go doing(conn)
		num++
		if num == 100000 {
			log.Println(num)
			time.Sleep(time.Second * 300)
			break
		}

		// time.Sleep(time.Second * 1)
	}
}
