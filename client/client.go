package client

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"fly_shell/config"
	"io"
	"bufio"
)



func FlyClient(username, port, addr, command string, out chan string){
	address :=  addr + ":" + port
	sshconfig := config.SshConfig(username)
	client, err := ssh.Dial("tcp", address, sshconfig)
	if err != nil{
		panic("Failed to dail ssh: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil{
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	stdout, err := session.StdoutPipe()
	if err != nil{
		fmt.Println(err.Error())
	}
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr
	//session.Stdin = os.Stdout
	go batch_out(stdout, out)
	err = session.RequestPty("xterm-256color", 80, 40, *modes())
	if err != nil{
		fmt.Println(err.Error())
	}
	session.Run(command)
}

func modes() *ssh.TerminalModes{
	mode := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	return &mode
}


func batch_out(stdout io.Reader, out chan string) {
	outreader := bufio.NewReader(stdout)
	for {
		line, err := outreader.ReadString('\n')
		if err != nil || io.EOF == err {
			if err != io.EOF {
				panic(fmt.Sprintf("faield to execute command: %s", err))
			}
			break
		}
		out <- line
	}
}