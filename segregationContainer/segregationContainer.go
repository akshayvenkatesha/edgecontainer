package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var shareIp string
var sourceShare string
var localShare string
var cloudShare string
var shareUserName string
var sharePassword string

const sourceShareMountPath = "/source"
const localShareMountPath = "/local"
const cloudShareMountPath = "/cloud"

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

	createInitFolders()
	mountShares()
	unMountShares()
}

func createInitFolders() {
	createFolder(sourceShareMountPath)
	createFolder(localShareMountPath)
	createFolder(cloudShareMountPath)
}

func createFolder(folderPath string) {
	err := os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		fmt.Errorf("failed to create folder" + folderPath)
	}
}

func mountShares() {
	mountShare(sourceShare, sourceShareMountPath)
	// mountShare(localShare, localShareMountPath)
	// mountShare(cloudShare, cloudShareMountPath)
}

func mountShare(sharePath string, mountPath string) {
	str := "mount -t cifs //" + shareIp + sharePath + " " + mountPath + " -o username=" + shareUserName + ",password=" + sharePassword
	fmt.Println(str)
	cmd := exec.Command("sh", "-c", str)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

func unMountShares() {
	unMountShare(sourceShareMountPath)
	// mountShare(localShareMountPath)
	// mountShare(cloudShareMountPath)
}

func unMountShare(mountPath string) {
	str := "umount " + mountPath
	fmt.Println(str)
	cmd := exec.Command("sh", "-c", str)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
