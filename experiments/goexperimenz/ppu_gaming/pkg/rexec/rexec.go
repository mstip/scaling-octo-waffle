package rexec

import (
	"errors"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

type Rexecer interface {
	Exec(host string, cmd string) (string, error)
}

type Rexec struct {
	sshConfig *ssh.ClientConfig
}

func NewRexec(pathToKeyFile string, username string) *Rexec {
	key, err := ioutil.ReadFile(pathToKeyFile)
	if err != nil {
		log.Fatal("unable to read private key: " + err.Error())
		return nil
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("unable to parse private ke " + err.Error())
		return nil
	}

	r := &Rexec{}

	r.sshConfig = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return r
}

func (r Rexec) Exec(host string, cmd string) (string, error) {
	client, err := ssh.Dial("tcp", host+":22", r.sshConfig)
	if err != nil {
		return "", errors.New("unable to connect " + err.Error())
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", errors.New("Failed to create session " + err.Error())
	}
	defer session.Close()

	out, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Print(err)
		return string(out), errors.New("cmd runtime error " + err.Error())
	}
	return string(out), nil
}
