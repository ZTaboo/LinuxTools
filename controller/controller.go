package controller

import (
	"fmt"
	"os/exec"
)

func Version(data string) {
	if data == "true" {
		cmd := exec.Command("lsb_release", "-a")
		con, _ := cmd.Output()
		fmt.Println(string(con))
	}
}

func ReplaceOsSource(data string) {
	if data == "true" {
		cmd := exec.Command("lsb_release", "-r")
		con, _ := cmd.Output()
		switch string(con[9 : len(con)-1]) {
		case "20.04":
			fmt.Println("this is 20.04")
		}
	}
}
