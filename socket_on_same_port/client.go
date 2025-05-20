package main

import (
	"fmt"
	"net"
	"time"
)

const (
	server   = "server"
	port     = "1234"
	interval = 10 * time.Second
)

func udpClient() {
	localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("UDP resolve local:", err)
		return
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", server+":"+port)
	if err != nil {
		fmt.Println("UDP resolve remote:", err)
		return
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("UDP dial:", err)
		return
	}
	defer conn.Close()

	for {
		msg := fmt.Sprintf("Hello UDP at %v \n", time.Now().Unix())
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("UDP send error:", err)
			return
		}
		fmt.Println("Sent UDP:", msg)
		time.Sleep(interval)
	}
}

func tcpClient() {
	localAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("TCP resolve local:", err)
		return
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", server+":"+port)
	if err != nil {
		fmt.Println("TCP resolve remote:", err)
		return
	}

	conn, err := net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		fmt.Println("TCP connect error:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Connected to %s from %s\n", remoteAddr, localAddr)

	for {
		msg := fmt.Sprintf("Hello TCP at %v \n", time.Now().Unix())
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("TCP send error:", err)
			return
		}
		fmt.Println("Sent TCP:", msg)
		time.Sleep(interval)
	}
}

func main() {
	go udpClient()
	go tcpClient()
	select {} // Block forever
}
