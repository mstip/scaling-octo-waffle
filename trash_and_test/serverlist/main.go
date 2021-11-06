package main

import (
	"io/ioutil"
	"log"
	"net"
	"sort"
	"strings"
	"time"
)

func serverCheck(hostname string) bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", hostname+":22", timeout)

	if err != nil {
		return false
	}
	return true
}

func main() {
	// read
	data, err := ioutil.ReadFile("./serverlist.txt")
	if err != nil {
		log.Fatal(err)
	}

	// split by newline
	list := strings.Split(string(data), "\n")
	serverlist := map[string]bool{}
	// make unique and add the domain as suffix
	for _, v := range list {
		if !strings.HasSuffix(v, "XXX") {
			v = v + "XXX"
		}
		serverlist[v] = true
	}

	// check if the servers exist by checking for ssh
	uniqueExistingServers := []string{}
	for k := range serverlist {
		if serverCheck(k) {
			uniqueExistingServers = append(uniqueExistingServers, k)
		}
	}

	// sort
	sort.Strings(uniqueExistingServers)
	// write result
	ioutil.WriteFile("./result.txt", []byte(strings.Join(uniqueExistingServers, "\n")), 666)
}
