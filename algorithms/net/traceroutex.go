package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: traceroute <destination>")
		return
	}

	destAddr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Error resolving destination address:", err)
		return
	}

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("Error creating ICMP packet listener:", err)
		return
	}
	defer conn.Close()

	for ttl := 1; ttl <= 30; ttl++ {
		conn.SetTTL(ttl)

		message := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  1,
				Data: []byte("traceroute"),
			},
		}

		messageBytes, err := message.Marshal(nil)
		if err != nil {
			fmt.Println("Error marshalling ICMP message:", err)
			return
		}

		start := net.Now()

		if _, err := conn.WriteTo(messageBytes, destAddr); err != nil {
			fmt.Println("Error sending ICMP message:", err)
			return
		}

		reply := make([]byte, 1500)
		n, _, err := conn.ReadFrom(reply)
		if err != nil {
			fmt.Println("Error receiving ICMP reply:", err)
			return
		}

		duration := net.Now().Sub(start)

		addr := "<unknown>"
		if src, err := net.ResolveIPAddr("ip", string(reply[:n])); err == nil {
			addr = src.String()
		}

		fmt.Printf("%d: %s (%s) %v\n", ttl, destAddr.String(), addr, duration)

		if addr == destAddr.IP.String() {
			break
		}
	}
}
