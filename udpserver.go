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
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:9002")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024*1024)
	// data := []byte{}
	rcv_len, _, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}

	if rcv_len <= 4 {
		fmt.Println("rcv_error, len <= 4")
		fmt.Println(rcv_len)
		return
	}

	fmt.Println(string(data[:rcv_len]))
	// length := int(binary.BigEndian.Uint32(data))
	// // in := string(data[:])
	// fmt.Println(length)

	// ret := make([]byte, rcv_len)
	// _, remoteAddr, err = conn.ReadFromUDP(content)
	// if err != nil {
	// 	fmt.Println("failed to read UDP msg because of ", err.Error())
	// 	return
	// }

	// fmt.Println(string(content[:]))

	// ret := make([]byte, length)
	//
	// copy(ret, data)
	// copy(ret[4:], content)

	/////
	// conn.WriteToUDP(ret, remoteAddr)
}

// func handleClient(conn *net.UDPConn) {
// 	// data := []byte{}
// 	data := make([]byte, 10000)
// 	oob := make([]byte, 10000)
// 	n, oobn, f, _, err := conn.ReadMsgUDP(data, oob)
// 	if err != nil {
// 		fmt.Println("failed to read UDP msg because of ", err.Error())
// 		return
// 	}
//
// 	fmt.Println("----------ret-----------")
// 	fmt.Println(n)
// 	fmt.Println(oobn)
// 	fmt.Println(f)
//
// 	return
// 	//
// 	// length := int(binary.BigEndian.Uint32(data))
// 	// // in := string(data[:])
// 	// fmt.Println(length)
// 	//
// 	// content := make([]byte, length)
// 	// _, remoteAddr, err = conn.ReadFromUDP(content)
// 	// if err != nil {
// 	// 	fmt.Println("failed to read UDP msg because of ", err.Error())
// 	// 	return
// 	// }
// 	//
// 	// // fmt.Println(string(content[:]))
// 	//
// 	// // ret := make([]byte, length)
// 	// //
// 	// // copy(ret, data)
// 	// // copy(ret[4:], content)
// 	//
// 	// /////
// 	// conn.WriteToUDP(content, remoteAddr)
// }
