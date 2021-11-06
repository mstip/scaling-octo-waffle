package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("accepted conn", conn)
		go func(conn net.Conn) {
			log.Println("handle conn local addr", conn.LocalAddr().String())
			log.Println("handle conn remote addr", conn.RemoteAddr().String())
			//b, err := io.ReadAll(conn)

			go func() {
				for {
					log.Println("send", "woop")

					_, err := conn.Write([]byte("woop\n"))

					//_, err = fmt.Fprint(conn, "woop\n")
					if err != nil {
						log.Fatal(err)
					}

					//if _, err := conn.Write([]byte("Hello, World!\n")); err != nil {
					//	log.Fatal(err)
					//}

					time.Sleep(3 * time.Second)
				}
			}()

			for {
				//b, err := bufio.NewReader(conn).ReadBytes(0)
				//b, err := bufio.NewReader(conn).ReadByte()
				b := make([]byte, 512)
				n, err := conn.Read(b)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("read", n, b)

			}

			//conn.Close()
			log.Println("end conn")
		}(conn)
	}
}
