package cell

import (
	"net"
	"os"
)

func GetIPv4() (ip net.IP) {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = ipv4
		}
	}
	return
}
