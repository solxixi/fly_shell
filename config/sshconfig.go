package config

import (
	"golang.org/x/crypto/ssh"
	"os/user"
	"io/ioutil"
	"fmt"
)

func SshConfig(username string) (config *ssh.ClientConfig){
	key, err := GetPubKey()
	if err != nil{
		return
	}
	config = &ssh.ClientConfig{
		User: username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	return config
}


func GetPubKey() (key ssh.Signer, err error) {
	userinfo, _ := user.Current()
	file := userinfo.HomeDir + "/.ssh/id_rsa"
	buf, err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Println("读取私钥错误")
		return
	}
	key ,err = ssh.ParsePrivateKey(buf)
	if err != nil{
		return
	}
	return
}