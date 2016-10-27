package main

import (
	"esmail/utils"
	"fmt"
	"log"

	"github.com/dgiagio/getpass"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	// VERSION of ESMail
	VERSION = "0.1"
	// ConfFile where save conf
	confPath = "./conf/esmail.json"
)

var (
// EmailPW = kingpin.Flag("pw", "your email password").Required().String()
)

// esc EasyMail Conf
var esc struct {
	Sender string `json:"sender"`
	ESHost string `json:"eshost"`
}

func init() {
	kingpin.Version(VERSION)
	kingpin.Parse()
}

func main() {
	err := utils.ParseJSON(&esc, confPath)
	if err != nil {
		panic(err)
	}
	log.Println(esc)

	pw, err := getpass.GetPassword("Password:")
	if err != nil {
		log.Println("GetPassword error =>", err)
	}
	fmt.Println("GetPassword => ", pw)
}
