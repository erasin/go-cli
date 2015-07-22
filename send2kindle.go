package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/gomail.v1"
)

type Config struct {
	Smtp   Smtp   `toml:"smtp"`
	From   string `toml:"from"`
	Kindle string `toml:"kindle"`
}

type Smtp struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Account  string `toml:"account"`
	Password string `toml:"password"`
}

func (smtp Smtp) Send(msg *gomail.Message) {
	mailer := gomail.NewMailer(smtp.Host, smtp.Account, smtp.Password, smtp.Port)
	if err := mailer.Send(msg); err != nil {
		panic(err)
	}
}

func getContent(url string) []byte {
	body, err := exec.Command("w3m", "-dump", url).Output()
	if err != nil {
		log.Fatal("w3m:", err)
	}
	return body
}

func main() {
	var config Config

	usr, _ := user.Current()
	configPath := path.Join(usr.HomeDir, ".send2kindle.toml")

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s url", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	segments := strings.Split(url, "/")

	msg := gomail.NewMessage()
	msg.SetHeader("From", config.From)
	msg.SetHeader("To", config.Kindle)
	msg.SetHeader("Subject", "your new article")
	msg.SetBody("text/html", "")
	msg.Attach(&gomail.File{
		Name:     segments[len(segments)-1] + ".txt",
		MimeType: "text/plain",
		Content:  getContent(url),
	})
	config.Smtp.Send(msg)

	fmt.Println("sent")
}
