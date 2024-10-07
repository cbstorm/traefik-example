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
	if port == 0 {
		log.Fatalf("Port is required")
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Could not listen due to an error: %v", err)
	}
	defer listener.Close()
	log.Printf("TCP server [%s] listening at %s:%d", sName, host, port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleConn(conn)
	}
}
func HandleConn(conn net.Conn) error {
	b := make([]byte, 1440)
	n, err := conn.Read(b)
	if err != nil {
		log.Printf("Could not data from request due to an error: %v", err)
		if _, err := conn.Write([]byte(err.Error())); err != nil {
			log.Fatalf("Could not write error message to client due to an error: %v", err)
		}
	}
	log.Printf("%s: %s", conn.RemoteAddr(), string(b))
	if _, err := conn.Write([]byte(fmt.Sprintf("Server %s received %d bytes", sName, n))); err != nil {
		log.Fatalf("Could not write error message to client due to an error: %v", err)
	}
	return conn.Close()
}
