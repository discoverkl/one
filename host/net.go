package host

import (
	"fmt"
	"net"
)

// PrimaryIP of current host.
func PrimaryIP() (string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifs {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.To4()
			case *net.IPAddr:
				ip = v.IP.To4()
			}
			if ip != nil && !ip.IsLoopback() {
				return ip.String(), nil
			}
		}
	}
	return "", fmt.Errorf("IP not found")
}
