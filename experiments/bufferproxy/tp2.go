package main

import (
	"errors"
	"log"
	"net"
	"time"
)

const reconnectRetries = 300
const reconnectTimeoutMs = 100
const listenTo = ":3000"
const redirectTo = "localhost:5672"

func remoteConnect(addr string) (net.Conn, error) {
	var remoteConn net.Conn
	var err error
	for n := 0; n < reconnectRetries; n++ {
		remoteConn, err = net.Dial("tcp", addr)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(reconnectTimeoutMs * time.Millisecond)
		if n == reconnectRetries-1 {
			log.Println("give up")
			return nil, errors.New("giveup")
		}
	}
	return remoteConn, nil
}

func proxyCopy(readFromConn net.Conn, writeToConn net.Conn) error {
	//log.Println("start proxy copy")
	b := make([]byte, 512)
	n, err := readFromConn.Read(b)
	if err != nil {
		log.Println("while read")
		return err
	}
	if n >= 512 {
		return errors.New("not implemented buffer resize")
	}
	//log.Println("read ", n, b[0:n], string(b[0:n]))
	n, err = writeToConn.Write(b[0:n])
	if err != nil {
		log.Println("while write")
		return err
	}
	//log.Println("write and end", n, b[0:n])
	return nil
}

func proxyCopyEndlessWrapper(readFromConn net.Conn, writeToConn net.Conn, prefix string) {
	for {
		err := proxyCopy(readFromConn, writeToConn)
		if err != nil {
			log.Println(prefix, "err", err)
			break
		}
	}
}

func handletwo(conn net.Conn) {
	remote, err := remoteConnect(redirectTo)
	if err != nil {
		log.Fatal(err)
	}

	go proxyCopyEndlessWrapper(remote, conn, "r-c")
	go proxyCopyEndlessWrapper(conn, remote, "c-r")
	time.Sleep(time.Second * 60)
}

func main() {
	ln, err := net.Listen("tcp", listenTo)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handletwo(conn)
	}
}
