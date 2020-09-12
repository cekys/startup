package main

import (
	"github.com/cekys/gopkg"
	"log"
	"os/exec"
)

const (
	//VNCServer VNC Server Path
	VNCServer = "C:/Users/Administrator/Desktop/Tools/UltraVNC/winvnc.exe "
)

func main() {
	cmd := exec.Command(VNCServer)
	mypkg.ShowArgs()
	//Run cmd without waiting for it to complete
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
