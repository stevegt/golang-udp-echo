package main

import (
	"fmt"
	"net"
	"os"
)

func ck(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func serve(port string) (done chan bool) {
	local_addr, err := net.ResolveUDPAddr("udp", port)
	ck(err)
	conn, err := net.ListenUDP("udp", local_addr)
	ck(err)
	buf := make([]byte, 65536)
	done = make(chan bool)

	go func() {
		defer conn.Close()
		for {
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			fmt.Println("Received ", string(buf[0:n]), " from ", addr)
			// rx <- buf[0:n]
			conn.WriteToUDP(buf[0:n], addr)
		}
	}()
	return
}

func main() {
	done := serve(":8981")
	<-done
}
