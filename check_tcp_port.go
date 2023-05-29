package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	TCP_TIMEOUT = "10s"
)

func Check_tcp_port(host string, port string) (string, error) {
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return "", fmt.Errorf("Invalid port %q: %s", port, err)
	}

	tcp_timeout_duration, _ := time.ParseDuration(TCP_TIMEOUT)
	ln, err := net.DialTimeout("tcp", host+":"+port, tcp_timeout_duration)

	if err != nil {
		return "", fmt.Errorf("Can't listen on %s:%q: %s", host, port, err)
	}

	err = ln.Close()
	if err != nil {
		return "", fmt.Errorf("Couldn't stop listening on %s:%q: %s", host, port, err)
	}

	return fmt.Sprintf("%s:%s is up!", host, port), nil
}
