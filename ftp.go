package main

import (
	"flag"
	"fmt"
	"pasvftp/ftp"
	// "github.com/molizz/pasvftp/ftp"
)

func main() {
	fmt.Println("pasvftp by moli")
	var remoteHost string
	flag.StringVar(&remoteHost, "h", "remoteHost", "remoteHost")
	flag.Parse()

	p := ftp.NewFtpProxy(2121, remoteHost, 21)
	_ = p.Start()
}
