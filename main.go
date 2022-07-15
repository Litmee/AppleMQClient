package main

import (
	"AppleMQClient/treaty"
	"encoding/json"
	"log"
	"net"
	"time"
)

type AppleMessage struct {
	// Sign: 0 means from the producer, 1 means from the cluster machine synchronization
	Sign int
	Body string
	Tag  string
}

func doing(c *net.TCPConn) {
	message := &AppleMessage{Sign: 0, Body: time.Now().String(), Tag: "ok"}

	marshal, _ := json.Marshal(message)

	encode, err := treaty.Encode(string(marshal))
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
	encode, _ := treaty.Encode("send")
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
