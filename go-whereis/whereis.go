package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ttacon/chalk"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("nothing!!!")
		os.Exit(0)
	}

	fc := os.Args[1:]

	for i := 0; i < len(fc); i++ {
		path, err := exec.LookPath(fc[i])
		if err != nil {
			fmt.Printf("%s%s \t%s-> %scan't find.%s\n", chalk.Red, fc[i], chalk.Reset, chalk.Red, chalk.Reset)
		} else {
			fmt.Printf("%s%s \t%s-> %s\n", chalk.Blue, fc[i], chalk.Reset, path)
		}
	}

}
