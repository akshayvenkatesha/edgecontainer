package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func printFlags() {
	fmt.Println()
	fmt.Println("Input Details:")
	fmt.Println("shareIp:" + shareIp)
	fmt.Println("sourceShare:" + sourceShare)
	fmt.Println("localShare:" + localShare)
	fmt.Println("cloudshare:" + cloudShare)
	fmt.Println("shareUserName:" + shareUserName)
	fmt.Println("sharePassword:" + sharePassword)
	fmt.Println()
}

func main() {
	initFlags()
	flag.Parse()
	printFlags()
	createInitFolders()
	mountShares()
	businessLogin()
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
	mountShare(localShare, localShareMountPath)
	mountShare(cloudShare, cloudShareMountPath)
	fmt.Println()
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

func businessLogin() {
	err := filepath.Walk(sourceShareMountPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() == true {
				return nil
			}
			shareToCopy := decideTheShare(info)
			fmt.Println(path, info.Size())
			fmt.Println(path, shareToCopy)
			copyFile(path, shareToCopy)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func copyFile(path string, shareToCopy string) {
	in, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func decideTheShare(info os.FileInfo) string {
	if info.Size() > 100000 {
		return cloudShareMountPath
	}
	return localShareMountPath
}

func unMountShares() {
	fmt.Println()
	unMountShare(sourceShareMountPath)
	unMountShare(localShareMountPath)
	unMountShare(cloudShareMountPath)
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
