package main

import (
	"flag"
	"fmt"
)

var shareIp string
var sourceShare string
var localShare string
var cloudShare string
var shareUserName string
var sharePassword string

func initFlags() {
	flag.StringVar(&shareIp, "shareIp", "127.0.0.1", "Ip of the machine where shares are hosted.")
	flag.StringVar(&sourceShare, "sourceShare", "/sourceShare", "Source share from where files should be read.")
	flag.StringVar(&localShare, "localShare", "/localShare", "Local share to where local files should be moved to.")
	flag.StringVar(&cloudShare, "cloudShare", "/cloudshare", "Cloud share to where cloud files should be moved to. ")
	flag.StringVar(&shareUserName, "shareUserName", "testuser", "Share user name which will be used to connect.")
	flag.StringVar(&sharePassword, "sharePassword", "Password", "Share password.")
}

func main() {
	initFlags()
	flag.Parse()
	fmt.Println(shareIp)
	fmt.Println(sourceShare)
	fmt.Println(localShare)
	fmt.Println(cloudShare)
	fmt.Println(shareUserName)
	fmt.Println(sharePassword)
}
