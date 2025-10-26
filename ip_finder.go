package main

import (
	"net"
)

func GetLocalIPs(l *Logger) []string {
	var ips []string

	ifaces, err := net.Interfaces()

	if err != nil {
		l.Error(err)
		return ips
	}

	for _, iface := range ifaces {

		// to skip inactive interfaces and loopback interfaces.
		if (iface.Flags & net.FlagUp) == 0 {
			l.Info(iface.Name + " [Down]")
			continue
		}
		if (iface.Flags & net.FlagLoopback) != 0 {
			l.Info(iface.Name + " [Loopback]")
			continue
		}
		addrs, err := iface.Addrs()

		if err != nil {
			l.Error(err)
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ip := ipNet.IP

			if ip.To4() != nil {
				l.Info(iface.Name+": "+ip.String() + " [UP]")
				ips = append(ips,iface.Name + ": " + ip.String())
			}
		}
	}
	return ips
}
