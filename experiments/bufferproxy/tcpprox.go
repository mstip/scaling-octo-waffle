//package main
//
//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"time"
//)
//
//func remoteConnect(addr string) (net.Conn, error) {
//	var remoteConn net.Conn
//	var err error
//	for n := 0; n < 300; n++ {
//		remoteConn, err = net.Dial("tcp", addr)
//		if err == nil {
//			break
//		}
//		log.Println(err)
//		time.Sleep(100 * time.Millisecond)
//		if n == 299 {
//			log.Println("give up")
//			return nil, errors.New("giveup")
//		}
//	}
//	return remoteConn, nil
//}
//
//func proxyCopy(readFromConn net.Conn, writeToConn net.Conn) error {
//	log.Println("start proxy copy")
//	b := make([]byte, 512)
//	n, err := readFromConn.Read(b)
//	if err != nil {
//		return err
//	}
//	if n >= 512 {
//		log.Fatal("not implemented buffer resize")
//	}
//	log.Println("read ", n, b[0:n], string(b[0:n]))
//	n, err = writeToConn.Write(b[0:n])
//	if err != nil {
//		return err
//	}
//	log.Println("write and end", n, b[0:n])
//	return nil
//}
//
//func handletwo(conn net.Conn) {
//	//remote, err := remoteConnect("localhost:3030")
//	remote, err := remoteConnect("localhost:5672") // rabbit
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//for {
//	log.Println("alive !")
//	go func() {
//		for {
//
//			err = proxyCopy(remote, conn)
//			if err != nil {
//				//log.Println("reconnect")
//				//remote, err = remoteConnect("localhost:5672") // rabbit
//				if err != nil {
//					log.Fatal(err)
//				}
//			}
//		}
//	}()
//	go func() {
//		for {
//			err = proxyCopy(conn, remote)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//		}
//	}()
//	time.Sleep(time.Second * 60)
//	//}
//
//}
//
//func handle(conn net.Conn) {
//	fmt.Println("new connection")
//
//	//var remoteConn net.Conn
//	//var err error
//	//for n := 0; n < 300; n++ {
//	//	remoteConn, err = net.Dial("tcp", "google.com:80")
//	//	remoteConn, err = net.Dial("tcp", "localhost:5672")
//	//remoteConn, err = net.Dial("tcp", "localhost:3030")
//	//if err == nil {
//	//	break
//	//}
//	//log.Println(err)
//	//time.Sleep(100 * time.Millisecond)
//	//if n == 299 {
//	//	log.Println("give up")
//	//	if err = conn.Close(); err != nil {
//	//		log.Println(err)
//	//	}
//	//	return
//	//}
//	//}
//
//	remoteConn, err := remoteConnect("localhost:3030")
//	if err != nil {
//		conn.Close()
//		log.Fatal(err)
//	}
//
//	go func() {
//		//n, err := io.Copy(remoteConn, conn)
//
//		b, err := bufio.NewReader(conn).ReadBytes('\n')
//		log.Println(b)
//		n, err := fmt.Fprint(remoteConn, b)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		log.Println("conn to remote", n, err)
//		log.Println("timeout", err.(net.Error).Timeout())
//		log.Println("temp", err.(net.Error).Temporary())
//		log.Println("eof", err == io.EOF)
//
//		log.Println("reconnect")
//		remoteConn, err = remoteConnect("localhost:3030")
//		if err != nil {
//			log.Fatal(err)
//		}
//		//n, err = io.Copy(remoteConn, conn)
//		log.Println("conn to remote", n, err)
//	}()
//	n, err := io.Copy(conn, remoteConn)
//
//	log.Println("remote to conn", n, err)
//
//	if err = remoteConn.Close(); err != nil {
//		log.Println(err)
//	}
//	if err = conn.Close(); err != nil {
//		log.Println(err)
//	}
//}
//
//func main() {
//	ln, err := net.Listen("tcp", ":3000")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			log.Fatal(err)
//		}
//		//go handle(conn)
//		go handletwo(conn)
//	}
//}
