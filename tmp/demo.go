package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

func main() {
	a, b := returnTwo() // Assignment count mismatch: 2 = 1
	fmt.Println(a, b)
}

func listenOnSpecificInterface(p *ipv4.PacketConn, iface *net.Interface) {
	mDNSLinkLocal := net.UDPAddr{IP: net.IPv4(224, 0, 0, 251)}
	if err := p.JoinGroup(iface, &mDNSLinkLocal); err != nil {
		log.Fatal(err)
	}
	defer p.LeaveGroup(iface, &mDNSLinkLocal)
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 1500)
	for {
		fmt.Printf("Waiting for data on interface: %v\n", iface.Name)
		n, cm, peer, err := p.ReadFrom(b)
		fmt.Printf("Received on interface %v: %v \n", iface.Name, string(b[:n]))
		if err != nil {
			log.Fatal(err)
		}
		if !cm.Dst.IsMulticast() || !cm.Dst.Equal(mDNSLinkLocal.IP) {
			continue
		}
		answers := []byte("FAKE-MDNS-ANSWERS") // fake mDNS answers, you need to implement this
		if _, err := p.WriteTo(answers, nil, peer); err != nil {
			log.Fatal(err)
		}
	}
}

func returnTwo() (int, int) {
	return 1, 2
}
