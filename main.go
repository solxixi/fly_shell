package main

import (
	"fly_shell/client"
	"flag"
	"os/user"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main()  {
	var w sync.WaitGroup
	out := make(chan string, 300)
	username := flag.String("u", "", "remote ssh user")
	port := flag.String("p", "22", "remote ssh port")
	addrlist := flag.String("h", "", "remote ip address")
	command := flag.String("c", "", "remote command")
	flag.Parse()
	userinfo, _ := user.Current()
	if *username == "" {
		*username = userinfo.Username
	}
	if *addrlist == "" || *command == "" {
		flag.Usage()
		fmt.Fprint(os.Stderr, "\n")
		os.Exit(1)
	}

	addr := strings.Split(*addrlist,",")
	for _, h := range addr{
		fmt.Println(h)
		w.Add(1)
		defer w.Done()
		go client.FlyClient(*username, *port, h, *command, out)
	}

	go func() {
		for output := range out {
			content := strings.Trim(output, "\r\n") + "\n"
			fmt.Printf(content)
		}
	}()

	w.Wait()
}