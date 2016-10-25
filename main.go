package main

import (
	"fmt"
	"log"

	"github.com/dgiagio/getpass"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	VERSION = "0.1"
)

var (
// EmailPW = kingpin.Flag("pw", "your email password").Required().String()
)

func init() {
	kingpin.Version(VERSION)
	kingpin.Parse()
}

func main() {
	var pw string
	pw, err := getpass.GetPassword("Password:")
	if err != nil {
		log.Println("GetPassword error =>", err)
	}
	fmt.Println("GetPassword => ", pw)
}
