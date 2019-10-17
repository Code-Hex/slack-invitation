package ip

import (
	"net"
	"strings"
)

// Normalize normalize ip string. if contains ":", this function returns strings before of ":".
func Normalize(host string) (string, error) {
	if strings.Contains(host, ":") {
		var err error
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return "", err
		}
	}
	return host, nil
}
