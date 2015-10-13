package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"code.google.com/p/rsc/qr"
	// "strconv"
	"time"
)

func main() {
	// myqr := qr.Code{0, 100, 30, 1}
	if len(os.Args) < 2 {
		fmt.Println("you have input text!")
		os.Exit(0)
	}

	txt := os.Args[1]
	myqr, _ := qr.Encode(txt, qr.H)
	png := myqr.PNG()
	// name := strconv.FormatInt(time.Now().Unix(), 10) + ".png"
	name := time.Now().Format("20060102150405") + ".png"
	fmt.Println("file: ", name)
	ioutil.WriteFile(name, png, 0600)
}
