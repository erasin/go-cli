package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"

	// "github.com/astaxie/beego/httplib"
	// "github.com/jpoehls/gophermail"
)

var (
	ipExp = regexp.MustCompile(`(\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})`)
)

func main() {
	ip, err := globalIp()
	fmt.Println(ip, err)
	lip, err := localIp()
	fmt.Println(lip, err)

}

// 获取广域网IP
func globalIp() (ip string, err error) {
	url_ip := "http://ddns.oray.com/checkip"
	resp, err := http.Get(url_ip)
	if err != nil {
		return ip, err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		_ = "breakpoint"
		if err != nil {
			return ip, err
		}
		ip = ipExp.FindString(string(body))
		return ip, nil
	}
}

// 获取局域网IP
func localIp() (ip string, err error) {
	conn, err := net.Dial("udp", "baidu.com:80")
	defer conn.Close()
	if err != nil {
		return ip, err
	}
	ip, _, _ = net.SplitHostPort(conn.LocalAddr().String())
	return ip, nil
}
