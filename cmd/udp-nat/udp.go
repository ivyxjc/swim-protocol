package main

import (
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	src := "192.168.32.198:2500"
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	}
	listener, err := cfg.ListenPacket(context.Background(), "udp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()

	sendUDP("192.168.32.198:1600", "TRIGGER")

	fmt.Printf("UDP server start and listening on %s.\n", src)
	for {
		buf := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(listener, addr, buf[:n])
	}

}

func serve(listener net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Printf("%s\t: %s\n", addr, buf)
	listener.WriteTo([]byte("message recived!\n"), addr)
}

func sendUDP(addr, msg string) (string, error) {
	fmt.Println(addr)
	fmt.Printf("send msg: %s\n", msg)

	LocalAddr, _ := net.ResolveUDPAddr("udp", "192.168.32.198:2500")
	RemoteEP := net.UDPAddr{IP: net.ParseIP("192.168.32.198"), Port: 1600}
	conn, err := net.DialUDP("udp", LocalAddr, &RemoteEP)
	if err != nil {
		log.Printf("%+v", err)
	}
	fmt.Println(conn)
	//conn,_:=net.DialUDP("udp",addr)
	// send to socket
	_, err = conn.Write([]byte(msg))

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
