package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//conn, err := net.Dial("tcp", "localhost:3030")
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("conn established")

	go func() {
		for {
			b, err := bufio.NewReader(conn).ReadBytes('\n')

			if err != nil {
				log.Fatal(err)
			}
			log.Println("read", b)
		}

	}()

	for {
		log.Println("send woop")
		_, err = fmt.Fprint(conn, "woop\n")
		if err != nil {
			log.Fatal(err)
		}

		//if _, err := conn.Write([]byte("Hello, World!\n")); err != nil {
		//	log.Fatal(err)
		//}

		time.Sleep(3 * time.Second)
	}
	log.Println("close conn")
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("conn closed - end")
}
