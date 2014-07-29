package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	service := "locahost:"
	for i := 0; i < 8000; i++ {
		// 1s
		_, err := net.DialTimeout("tcp", service+strconv.Itoa(i), 1000000000)
		if err == nil {
			fmt.Println("port: ", i)
		}
	}
}
