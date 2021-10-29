package main

import (
	"fmt"
	"github.com/libp2p/go-reuseport"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := reuseport.ListenPacket("udp", "0.0.0.0:2500")
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	fmt.Println(l)
	defer l.Close()

	sendUDP("", "TRIGGER")
	sendUDP("", "TRIGGER")

	for {
		fmt.Println("++++++")
		buf := make([]byte, 1024)
		n, addr, err := l.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(l, addr, buf[:n])
	}

}

func sendUDP(addr, msg string) (string, error) {
	fmt.Println(addr)
	fmt.Printf("send msg: %s\n", msg)

	conn, err := reuseport.Dial("udp", "192.168.32.198:2500", "129.204.55.14:1600")
	//conn, err := reuseport.Dial("udp", "192.168.32.198:2500", "192.168.32.198:1600")
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

func serve(listener net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Printf("%s\t: %s\n", addr, buf)
	listener.WriteTo([]byte("message recived!\n"), addr)
}
