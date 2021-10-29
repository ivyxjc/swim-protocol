package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	src := "0.0.0.0:1600"
	listener, err := net.ListenPacket("udp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()
	fmt.Printf("UDP server start and listening on %s.\n", src)

	for {
		buf := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buf)
		if err != nil {
			continue
		}
		str := string(buf[:])
		fmt.Printf("str is [%s]\n", str)
		if strings.LastIndex(str, "TRIGGER") >= 0 {
			fmt.Println("???????????????????")
			fmt.Println(addr.String())
			sendUDP(addr.String(), "OK")
		}
		go serve(listener, addr, buf[:n])
	}

}

func serve(listener net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Printf("%s\t: %s\n", addr, buf)
	listener.WriteTo([]byte("message recived!\n"), addr)
}

func sendUDP(addr, msg string) (string, error) {
	conn, _ := net.Dial("udp", addr)
	// send to socket
	_, err := conn.Write([]byte(msg))
	// listen for reply
	bs := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
