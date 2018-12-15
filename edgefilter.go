package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	//ip := "192.168.1.9"
	createFolder("/source")
	createFolder("/cloud")
	createFolder("/local")

	cmd := exec.Command("sh", "-c", "mount -t cifs //192.168.1.9/share /source -o username=hcstestuser,password=StorSim1 ")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	files, err := ioutil.ReadDir("/source")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	cmd1 := exec.Command("sh", "-c", "umount /source")
	_, err1 := cmd1.CombinedOutput()
	if err1 != nil {
		log.Fatal(err)
	}
}

func createFolder(folderPath string) {
	err := os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		fmt.Errorf("failed to create folder" + folderPath)
	}
}
