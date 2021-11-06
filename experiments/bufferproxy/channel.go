package main

import "time"

func main() {
	blocker := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		blocker <- true
	}()
	 <-blocker
}
