package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var host string
var port int
var t bool

func main() {
	flag.StringVar(&host, "host", "", "Host")
	flag.IntVar(&port, "port", 0, "Port")
	flag.BoolVar(&t, "t", false, "TCP if specified, else UDP")
	flag.Parse()
	if port == 0 {
		log.Fatalf("Port is required")
	}
	if t {
		if err := tcp_client(); err != nil {
			log.Fatalf("Error occurred %v", err)
		}
		os.Exit(0)
	}
	if err := udp_client(); err != nil {
		log.Fatalf("Error occurred %v", err)
	}
	os.Exit(0)
}

func tcp_client() error {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	})
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("test tcp")); err != nil {
		return err
	}
	return nil
}

func udp_client() error {
	udp_addr := &net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}
	conn, err := net.DialUDP("udp", nil, udp_addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	n, err := conn.Write([]byte("test upd"))
	if err != nil {
		return err
	}
	log.Printf("Written to server [%s] %d bytes", conn.RemoteAddr(), n)
	return nil
}
