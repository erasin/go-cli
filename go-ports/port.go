package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	service := "192.168.199.184"
	for port := 10; port < 81; port++ {
		_, err := net.DialTimeout("tcp", service+strconv.Itoa(port), time.Second*3)
		if err == nil {
			fmt.Println(port)
		}
	}
}
