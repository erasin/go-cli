package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGlobalIp(t *testing.T) {
	Convey("本机广域网的IP检测", t, func() {
		ip, _ := globalIp()
		So(ip, ShouldNotBeBlank)
	})
}

func TestLocalIp(t *testing.T) {
	Convey("本机局域网的IP检测", t, func() {
		ip, _ := localIp()
		So(ip, ShouldNotBeBlank)
	})
}
