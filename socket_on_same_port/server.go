package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	host = "0.0.0.0"
	port = "1234"
)

func udpListener() {
	addr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		fmt.Println("UDP resolve error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("UDP listen error:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP listener up on %s:%s\n", host, port)

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("UDP read error:", err)
			continue
		}
		fmt.Printf("UDP/%s from %s\n", string(buf[:n]), remoteAddr)
	}
}

func tcpListener() {
	ln, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("TCP listen error:", err)
		return
	}
	defer ln.Close()

	fmt.Printf("TCP listener up on %s:%s \n", host, port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
			continue
		}
		go handleTCPClient(conn)
	}
}

func handleTCPClient(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("TCP/%s from %s \n", msg, addr)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("TCP read error:", err)
	}
}

func main() {
	go udpListener()
	go tcpListener()

	select {} // Block forever
}
