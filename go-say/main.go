package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	if _, err := exec.LookPath("say"); err != nil {
		fmt.Println("[Err]use say of osx!")
		os.Exit(1)
	}

	str := getflag()

	// fmt.Println(str)
	Say(str)

}

// 获取pipe
func getflag() (str string) {

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {

		if len(os.Args) == 1 {
			fmt.Printf("[Err] Usage: %s content", os.Args[0])
			os.Exit(1)
		}

		for _, s := range os.Args[1:] {
			str = str + s + " "
		}

	} else if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		// match(pattern, reader)

		for {
			input, err := reader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}
			str = str + string(input)
		}
	}

	return
}
