package main

import (
	"net"
	"time"
)

func ChkNetConn() {

	host := "inncu.cn"
	port := "80"
	timeout := time.Duration(5) * time.Second

	_, err := net.DialTimeout("tcp", (host + ":" + port), timeout)

	ChkErr(err)
}
