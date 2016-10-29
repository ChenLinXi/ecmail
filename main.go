package main

import (
	"errors"
	"log"
	"net/smtp"
	"strings"

	"github.com/dgiagio/getpass"
	utils "github.com/hackez/hzutils"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	// VERSION of ESMail
	VERSION = "0.1"
	// ConfFile where save conf
	confPath = "./conf/esmail.json"
)

// esc EasyMail Conf
var esc struct {
	User   string `json:"user"`
	ESHost string `json:"eshost"`
}

func init() {
	kingpin.Version(VERSION)
	kingpin.Parse()
}

func main() {
	err := utils.ParseJSONFromFile(confPath, &esc)
	if err != nil {
		panic(err)
	}
	log.Println("esmail conf => ", esc)

	pw, err := getpass.GetPassword("enter password:")
	if err != nil {
		log.Println("GetPassword error =>", err)
	}

	to := "xxx@qq.com;xxx@qq.com" // <<<==== enter the email address you want to send to.
	subject := "邮件服务器发送测试"
	body := `<html><head>Welcome to ESMail</head><body><h3>Have a <a href="http://www.github.com/HackeZ/ESMail">try</a> :)</h3></body></html>`
	log.Println("Sending ESMail...")
	err = sendESMail(esc.User, pw, esc.ESHost, to, subject, body, "html")
	if err != nil {
		log.Println("Sending Error =>", err)
		return
	}
	log.Println("ESMail Already Sent")
}

// Send only for service
// Params
//  - password	    login smtp server password
//  - to			example@qq.com;example1@163.com;example2@sina.com.cn;...
//  - subject		the subject of mail
//  - body			the content of mail
//  - mailType      mail type :html or text
func Send(pw, to, subject, body, mailType string) error {
	log.Println("Sending ESMail...")
	defer log.Println("ESMail Already Sent")

	err := utils.ParseJSONFromFile(confPath, &esc)
	if err != nil {
		return err
	}

	return sendESMail(esc.User, pw, esc.ESHost, to, subject, body, mailType)
}

// sendESMail send mail by stmp protocol
// Params
// 	- user  		example@example.com login smtp server user
//  - password	    login smtp server password
//  - host			smtp.example.com:port   smtp.qq.com:25
//  - to			example@qq.com;example1@163.com;example2@sina.com.cn;...
//  - subject		the subject of mail
//  - body			the content of mail
//  - mailType      mail type :html or text
func sendESMail(user, pw, host, to, subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	sendToWho := strings.Split(to, ";")

	auth := smtp.PlainAuth("", user, pw, hp[0])
	var contentType string
	switch mailType {
	case "html", "HTML", "Html":
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	case "text", "TEXT", "Text":
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	default:
		return errors.New("not support mail type, only for html and text")
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	return smtp.SendMail(host, auth, user, sendToWho, msg)
}
