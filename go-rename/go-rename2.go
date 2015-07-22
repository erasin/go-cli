package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	import_dir := "import"
	export_dir := "export"
	rename := map[string]string{
		"from": "to",
		"aaa":  "bbb",
	}

	if err := os.RemoveAll(export_dir); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(export_dir, 0777); err != nil {
		fmt.Println(err)
	}

	for k, v := range rename {
		s := strings.Split(v, "/")

		if len(s) >= 2 {
			t := export_dir + "/" + s[len(s)-2]
			_, err := os.Stat(t)
			if err != nil {
				if err := os.MkdirAll(t, 0777); err != nil {
					fmt.Println(err)
				}
			}
		}

		f := import_dir + "/" + k
		t := export_dir + "/" + v

		_, err := os.Stat(f)
		if err == nil {
			if err := os.Rename(f, t); err != nil {
				fmt.Println(err)
			}
		}
	}
}
