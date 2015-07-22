package main

import (
	"fmt"
	"io"
	// "io/ioutil"

	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	iconv "github.com/djimenez/iconv-go"
)

var log logs.BeeLogger

func main() {
	url := "http://www.wenku8.cn/novel/0/4/index.htm"

	index(url)
}

// 首页解析
func index(url string) {
	doc, _ := goquery.NewDocument(url)
	Title := doc.Find("#title").Text()
	Author := doc.Find("#info").Text()

	Title, _ = iconv.ConvertString(Title, "gbk2312", "utf-8")
	Author, _ = iconv.ConvertString(Author, "gb2312", "utf-8")

	fmt.Printf("Title: %s ", Title)
	fmt.Printf("\nAuthor: %s ", Author)

	fmt.Printf("\n---------")

	doc.Find("table.css td").Each(func(_ int, s *goquery.Selection) {
		Txt := s.Text()
		var Link string
		if s.HasClass("ccss") {
			// Link = s.Children("a").First().Attr("href")
			fmt.Printf("%v", s.Find("a").First().Attr("href"))
		} else {
			Link = "#"
		}
		Txt, _ = iconv.ConvertString(Txt, "gb2312", "utf-8")
		fmt.Printf("\n * [%s](%s)", Txt, Link)

	})

	// ioutil.WriteFile("test.txt", []byte(Title), 0600)
}

func g2u(s string) string {
	s, _ = iconv.ConvertString(s, "gbk2312", "utf-8")
	return s
}

func down(url string) {
	// log.Println("Now is read the page that's url is " + url)
	doc, _ := goquery.NewDocument(url)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		name_t := strings.Split(src, "/")
		name := name_t[len(name_t)-1]
		if FileExist("imgs/" + name) {
			log.Println(name + " is all exist.")
		} else {
			log.Println(name + " :loadding...")
			res, _ := http.Get(src)
			file, _ := os.Create("imgs/" + name)
			io.Copy(file, res.Body)
			log.Println(name + " over.")
		}
	})
}

// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
	// if err == nil {
	// 	if fi.Size() == 0 {
	// 		return false
	// 	}
	// }
	// return os.IsExist(err)
}
