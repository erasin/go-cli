package main

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCharIs(t *testing.T) {
	t.Parallel()

	convey.Convey("检测英文", t, func() {
		convey.So(charIS("english"), convey.ShouldEqual, EN)
	})

	convey.Convey("检测中文", t, func() {
		convey.So(charIS("中文内容"), convey.ShouldEqual, ZH)
	})

	convey.Convey("检测日文", t, func() {
		convey.So(charIS("坂本なほ"), convey.ShouldEqual, JP)
	})
}
