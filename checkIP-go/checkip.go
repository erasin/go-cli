package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	// "github.com/astaxie/beego/httplib"
	// "github.com/jpoehls/gophermail"
)

var (
	ipExp = regexp.MustCompile(`(\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})`)
)

func main() {
	url_ip := "http://ddns.oray.com/checkip"
	// req := httplib.Get(url_ip)
	// fmt.Println(req.String())

	var mailSubject string
	var mailBody string

	resp, err := http.Get(url_ip)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		ip := ipExp.FindString(string(body))
		fmt.Println(ip)
		mailSubject = "PI IP: " + ip
		mailBody = " <html> <body> <h3> " + ip + " </h3> </body> </html>"
	}

	fmt.Println("send email")
	// err = domail(mailUser, mailPassword, mailHost, mailTo, mailSubject, mailBody, "html")
	err = domail(mailTo, mailSubject, mailBody)

	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
	// time.Sleep(time.Second * 30)

}
