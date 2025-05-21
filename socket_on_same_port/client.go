package main

import (
	"fmt"
	"net"
	"time"
)

const (
	server   = "172.25.0.12"
	port     = "1234"
	interval = 10 * time.Second
)

func getIPsByInterface(name string) ([]string, error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // skip non-IPv4
		}
		ips = append(ips, ip.String())
	}
	return ips, nil
}

func udpClient(ip string) {
	localAddr, err := net.ResolveUDPAddr("udp", ip+":"+port)
	if err != nil {
		fmt.Printf("UDP resolve local (%s): %v\n", ip, err)
		return
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", server+":"+port)
	if err != nil {
		fmt.Printf("UDP resolve remote: %v\n", err)
		return
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		fmt.Printf("UDP dial error from %s: %v\n", ip, err)
		return
	}
	defer conn.Close()

	fmt.Printf("UDP connected from %s to %s\n", ip, remoteAddr)

	for {
		msg := fmt.Sprintf("Hello UDP from %s at %v", ip, time.Now().Unix())
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("UDP send error from %s: %v\n", ip, err)
			return
		}
		fmt.Println("Sent UDP:", msg)
		time.Sleep(interval)
	}
}

func tcpClient(ip string) {
	localAddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		fmt.Printf("TCP resolve local (%s): %v\n", ip, err)
		return
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", server+":"+port)
	if err != nil {
		fmt.Printf("TCP resolve remote: %v\n", err)
		return
	}

	conn, err := net.DialTCP("tcp", localAddr, remoteAddr)
	if err != nil {
		fmt.Printf("TCP connect error from %s: %v\n", ip, err)
		return
	}
	defer conn.Close()

	fmt.Printf("TCP connected from %s to %s\n", ip, remoteAddr)

	for {
		msg := fmt.Sprintf("Hello TCP from %s at %v\n", ip, time.Now().Unix())
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("TCP send error from %s: %v\n", ip, err)
			return
		}
		fmt.Println("Sent TCP:", msg)
		time.Sleep(interval)
	}
}

func main() {
	ips, err := getIPsByInterface("eth0")
	if err != nil {
		fmt.Println("Error getting IPs:", err)
		return
	}

	if len(ips) == 0 {
		fmt.Println("No IPv4 addresses found on eth0")
		return
	}

	for _, ip := range ips {
		go udpClient(ip)
		go tcpClient(ip)
	}

	select {} // block forever
}
