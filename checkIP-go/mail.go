package main

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

var (
	mailUser     = "arata@aliyun.com"
	mailPassword = "ling86dian"
	mailHost     = "smtp.aliyun.com"
	mailTo       = "erasinoo@gmail.com;"
	mailport     = ":465"
)

func domail(tomail string, title string, body string) (err error) {
	to, err := mail.ParseAddress(tomail)
	if err != nil {
		return
	}

	auth := smtp.PlainAuth(
		"",
		mailUser,
		mailPassword,
		mailHost,
	)

	header := make(map[string]string)
	header["From"] = mailUser
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err = smtp.SendMail(
		mailHost+mailport,
		auth,
		mailTo,
		[]string{to.Address},
		[]byte(message),
	)
	return err
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{
		Address: str,
		Name:    "",
	}

	return strings.Trim(addr.String(), " <>")
}
