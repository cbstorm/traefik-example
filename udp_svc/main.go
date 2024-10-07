package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var host string
var port int
var sName string

func main() {
	flag.StringVar(&host, "host", "", "Host")
	flag.IntVar(&port, "port", 0, "Port")
	flag.StringVar(&sName, "s_name", "", "Service name")
	flag.Parse()
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	log.Printf("UDP server listening at %s\n", addr.String())
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConn(conn, n, addr)
	}
}

func HandleConn(conn *net.UDPConn, n int, addr *net.UDPAddr) error {
	rmsg := fmt.Sprintf("Server [%s] received [%d bytes] from [%s]", sName, n, addr)
	fmt.Println(rmsg)
	if _, err := conn.WriteToUDP([]byte(rmsg), addr); err != nil {
		log.Printf("Could write to UDP [%s] due to an error: %v", addr.String(), err)
		return err
	}
	return nil
}
