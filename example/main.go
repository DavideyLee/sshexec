package main

import (
	"log"
	"sshexec"
	"time"
)

func main() {
	sshExecAgent := sshexec.SSHExecAgent{}
	sshExecAgent.Worker = 10
	sshExecAgent.TimeOut = time.Duration(120) * time.Second
	s, err := sshExecAgent.SftpHostByKey([]string{"192.168.227.128:22","172.16.1.128:22"}, "root", "example/main.go", "test.log")
	log.Println("res:",s)
	log.Println("err:",err)

	s1, err1 := sshExecAgent.SshHostByKey([]string{"192.168.227.128:22"}, "root", "ls -al")
	log.Println("res:",s1)
	log.Println("err:",err1)
	for {
		time.Sleep(1000)
	}
}
