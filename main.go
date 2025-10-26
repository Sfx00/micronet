package main

import (
	"flag"
	"fmt"
)

func main() {
	verbose := flag.Bool("v", false, "")

	flag.Parse()

	l := NewLogger(*verbose)
	l.Info("Scanning network interfaces...")
	ips := GetLocalIPs(l)

	if len(ips) == 0 {
		fmt.Println("[No IPs found!]")
		return
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
