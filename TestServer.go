package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received HTTP request from %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "Hello, World!")
	})

	go func() {
		// Listen for incoming UDP packets
		conn, err := net.ListenPacket("udp", ":12345")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		for {
			buf := make([]byte, 1024)
			n, addr, err := conn.ReadFrom(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Received UDP packet from %s: %s\n", addr.String(), string(buf[:n]))
		}
	}()

	go func() {
		// Listen for incoming TCP connections
		ln, err := net.Listen("tcp", ":12345")
		if err != nil {
			log.Fatal(err)
		}
		defer ln.Close()

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Received TCP connection from %s\n", conn.RemoteAddr().String())

			go func(conn net.Conn) {
				defer conn.Close()

				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Printf("Error reading from TCP connection: %s\n", err.Error())
					return
				}
				log.Printf("Received %d bytes over TCP from %s: %s\n", n, conn.RemoteAddr().String(), string(buf[:n]))
			}(conn)
		}
	}()

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
