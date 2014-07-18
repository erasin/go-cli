package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	mux := http.NewServeMux()

	// 静态文件 os 绝对路径
	wd, err := os.Getwd() // 当前路径
	if err != nil {
		log.Fatal(err)
	}

	// 前缀去除
	// 列出dir
	mux.Handle("/",
		http.StripPrefix("/",
			http.FileServer(http.Dir(wd))))

	log.Printf("Start Port: http://127.0.0.1:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}

}
