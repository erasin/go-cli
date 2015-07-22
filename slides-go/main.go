package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/present"
)

// github.com/sinmetal/slide/src/present
func main() {
	docFile := "test.slide"
	doc, err := parse(docFile, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doc)
}

func parse(name string, mode present.ParseMode) (*present.Doc, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return present.Parse(f, name, 0)
}
