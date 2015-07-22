package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	// "regexp"
	// "strconv"
	// "time"
)

const (
	Upload_Dir = "./"
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
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd))))
	mux.HandleFunc("/upload", Upload)
	mux.HandleFunc("/ip", IpShow)

	log.Printf("Start Port: http://127.0.0.1:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}

}

func Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		temp1 := `<!DOCTYPE html><html>
<head>
    <title>{{.}}</title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="http://cdn.staticfile.org/twitter-bootstrap/3.3.0/css/bootstrap.min.css">
	<!-- <link rel="stylesheet" href="http://cdn.staticfile.org/twitter-bootstrap/3.3.0/css/bootstrap-theme.min.css"> -->
	<link rel="stylesheet" href="http://cdn.staticfile.org/bootstrap-material-design/0.1.5/css/ripples.min.css">
	<link rel="stylesheet" href="http://cdn.staticfile.org/bootstrap-material-design/0.1.5/css/material-wfont.min.css">
	<link rel="stylesheet" href="http://cdn.staticfile.org/bootstrap-material-design/0.1.5/css/material.min.css">
</head>
<body>
<div class="container-fluid">
	<form enctype="multipart/form-data" action="/upload" method="post">
	<input type="file" name="uploadfile" multiple="">
	<input class='btn' type="submit" value="submit" />
	</form>
</div>
</body>
</html>
`
		// 创建一个 template
		t := template.New("Person Info")
		// 解析模板
		t, _ = t.Parse(temp1)
		t.Execute(w, "上传文件")
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Fprintf(w, "%v", "上传错误")
			return
		}
		fileext := filepath.Ext(handler.Filename)
		if check(fileext) == false {
			fmt.Fprintf(w, "%v", "不允许的上传类型")
			return
		}
		// filename := strconv.FormatInt(time.Now().Unix(), 10) + fileext
		filename := handler.Filename
		f, _ := os.OpenFile(Upload_Dir+filename, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "%v", "上传失败")
			return
		}
		filedir, _ := filepath.Abs(Upload_Dir + filename)
		// r.Header.Set("Content-type", "text/html")
		fmt.Fprintf(w, "%v", filename+"上传完成,服务器地址:"+filedir)
	}
}

func check(name string) bool {
	ext := []string{".exe"}

	for _, v := range ext {
		if v == name {
			return false
		}
	}
	return true
}

func IpShow(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "{\"IP\":\"%s\"}", ip)
}
