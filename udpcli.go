package main

import (
	// "encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	// var buf = make([]byte, 4)
	// binary.BigEndian.PutUint32(buf, uint32(10))
	//
	// // _, err = conn.Write([]byte("hello world"))
	// _, err = conn.Write(buf)
	// if err != nil {
	// 	fmt.Println("failed:", err)
	// 	os.Exit(1)
	// }

	content := "1234567890"
	_, err = conn.Write([]byte(content))
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}

	// data := make([]byte, 4+len(content))
	// _, err = conn.Read(data)
	// if err != nil {
	// 	fmt.Println("failed to read UDP msg because of ", err)
	// 	os.Exit(1)
	// }

	// // length := int(binary.BigEndian.Uint32(data))

	// fmt.Println(string(data[:]))
	// fmt.Println(length)
	os.Exit(0)
}
