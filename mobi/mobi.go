package main

import (
	"github.com/codeskyblue/go-sh"
	"github.com/russross/blackfriday"

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	// "os/exec"
	"regexp"
)

func main() {

	var filename string // 参数文件名
	var Tmp string      // 临时html文件
	var Mobi string     // mobi名称
	var Title string
	var Author string
	var Comment string
	var Lang string

	flag.StringVar(&Title, "t", "", "标题")
	flag.StringVar(&Author, "a", "", "作者")
	flag.StringVar(&filename, "f", "", "文件")
	flag.StringVar(&Comment, "c", "", "简介")
	flag.StringVar(&Lang, "l", "zh-CN", "语言")
	flag.Parse()

	if filename == "" {
		fmt.Println(" -f 输入文件路径")
		os.Exit(-1)
	}

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("文件不存在，检查路径")
		os.Exit(-1)
	}

	fmt.Printf("Title: %s\nAuthor: %s \n", Title, Author)
	// TODO: 添加从文件中读取

	Tmp = info.Name() + ".html"
	Mobi = info.Name() + ".mobi"

	fmt.Println("读取文件", filename)

	b, _ := ioutil.ReadFile(filename)
	re := blackfriday.HtmlRenderer(1, "title", "")
	Md := blackfriday.Markdown(b, re, blackfriday.EXTENSION_TABLES+blackfriday.EXTENSION_FENCED_CODE)

	fmt.Println("写入HTML")
	// 创建历史文件
	tmpfile, _ := os.Create(Tmp)
	defer os.Remove(Tmp)
	defer fmt.Println("删除临时文件HTML")
	defer tmpfile.Close()

	tmpfile.WriteString(fmt.Sprintf("<html><head><meta http-equiv='content-language' content='zh-CN' /><meta http-equiv='Content-type' content='text/html; charset=utf-8'><meta name='Author' content='%s'><title>%s</title></head><body>%s</body></html>", Author, Title, regexp.MustCompile(`\n`).ReplaceAllString(string(Md), "")))

	fmt.Println("创建MOBI")

	fmt.Print(fmt.Sprintf("ebook-convert %s %s --authors %s --comments '%s' --level1-toc '//h:h1' --level2-toc '//h:h2' --language '%s'\n", Tmp, Mobi, Author, Comment, Lang))

	sh.Command("ebook-convert", Tmp, Mobi, "--authors", Author, "--comments", Comment, "--level1-toc", "//h:h1", "--level2-toc", "//h:h2", "--language", Lang).Run()

	fmt.Println("结束")
}
