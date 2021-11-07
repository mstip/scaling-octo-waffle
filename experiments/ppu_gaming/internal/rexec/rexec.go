package rexec

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

var sshConfig *ssh.ClientConfig

func init() {
	key, err := ioutil.ReadFile("/home/dev/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	sshConfig = &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func Rexec(host string, cmd string) (string, error) {
	client, err := ssh.Dial("tcp", host+":22", sshConfig)
	if err != nil {
		log.Printf("unable to connect: %v", err)
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Print("Failed to create session: ", err)
		return "", err
	}
	defer session.Close()

	out, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Print(err)
		return string(out), err
	}
	return string(out), nil
}
